import 'jquery'
import 'bootstrap'

import './style.scss'

import storage from './storage.js'

let now = () => {
  let d = new Date()
  return d.getTime()
}

let stopRefresh = () => {
  clearInterval(window.ticker)
  $('#focusedInput').select()
  return false
}

let restartRefresh = () => {
  window.lastLoad = now()
  window.ticker = setInterval(tick, window.tickerInterval)
}

let setProgress = (perc) => {
  $('.progress-bar').css('width', `${perc}%`)
}

let loadPassword = () => {
  let options = loadOptions()
  $.get(`/v1/getPassword?length=${options.passwordLength}&special=${options.useSpecial}&xkcd=${options.useXKCD}`, (data) => {
    $('#focusedInput').val(data)
    window.lastLoad = now()
  })
}

let saveOptions = () => {
  let options = {
    passwordLength: $('#passwordLengthOption').val(),
    useSpecial: $('#useSpecialOption')[0].checked,
    useXKCD: $('#useXKCDOption')[0].checked,
  }

  storage.set('SecurePasswordOptions', options)
  $('#settingsModal').modal('hide')

  loadPassword()
}

let loadOptions = () => {
  let options = storage.get('SecurePasswordOptions')
  if (!options) {
    options = {
      passwordLength: 20,
      useSpecial: false,
      useXKCD: false,
    }
  }
  $('#passwordLengthOption').val(options.passwordLength)
  $('#useSpecialOption')[0].checked = options.useSpecial
  $('#useXKCDOption')[0].checked = options.useXKCD

  return options
}

let tick = () => {
  let diff = now() - window.lastLoad
  let perc = (window.refreshPassword - diff) / window.refreshPassword * 100
  setProgress(perc)
  if (diff >= window.refreshPassword) {
    loadPassword()
  }
}

$(() => {
  window.lastLoad = now()
  window.refreshPassword = 30000
  window.tickerInterval = 200

  window.ticker = setInterval(tick, window.tickerInterval)
  $('#focusedInput').bind('click', stopRefresh)
  $('#focusedInput').bind('blur', restartRefresh)
  $('#optionSave').bind('click', saveOptions)

  loadOptions()
  loadPassword()
})
