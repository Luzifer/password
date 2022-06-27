// Libraries
import $ from 'jquery'
import 'bootstrap'

// Styles
import './style.scss'

// FontAwesome 5
import {
  dom,
  library,
} from '@fortawesome/fontawesome-svg-core'
import {
  faCog,
} from '@fortawesome/free-solid-svg-icons'

library.add(faCog)
dom.watch()

// Application code
import storage from './storage.js'

const now = () => {
  const d = new Date()
  return d.getTime()
}

const stopRefresh = () => {
  clearInterval(window.ticker)
  $('#focusedInput').select()
  return false
}

const restartRefresh = () => {
  window.lastLoad = now()
  window.ticker = setInterval(tick, window.tickerInterval)
}

const setProgress = perc => {
  $('.progress-bar').css('width', `${perc}%`)
}

const loadPassword = () => {
  const options = loadOptions()
  $.get(`/v1/getPassword?length=${options.passwordLength}&special=${options.useSpecial}&xkcd=${options.useXKCD}&separator=${options.xkcdSeparator}`, data => {
    $('#focusedInput').val(data)
    window.lastLoad = now()
  })
}

const saveOptions = () => {
  const options = {
    passwordLength: $('#passwordLengthOption').val(),
    useSpecial: $('#useSpecialOption')[0].checked,
    useXKCD: $('#useXKCDOption')[0].checked,
    xkcdSeparator: $('#xkcdSeparator').val(),
  }

  storage.set('SecurePasswordOptions', options)
  $('#settingsModal').modal('hide')

  loadPassword()
}

const loadOptions = () => {
  let options = storage.get('SecurePasswordOptions')
  if (!options) {
    options = {
      passwordLength: 20,
      useSpecial: false,
      useXKCD: false,
      xkcdSeparator: '',
    }
  }
  $('#passwordLengthOption').val(options.passwordLength || 20)
  $('#useSpecialOption')[0].checked = options.useSpecial
  $('#useXKCDOption')[0].checked = options.useXKCD
  $('#xkcdSeparator').val(options.xkcdSeparator || '')

  return options
}

const tick = () => {
  const diff = now() - window.lastLoad
  const perc = (window.refreshPassword - diff) / window.refreshPassword * 100
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
