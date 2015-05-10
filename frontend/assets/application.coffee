if window.Storage and window.JSON
  window.$storage = (key) ->
    set: (value) ->
      localStorage.setItem(key, JSON.stringify(value))
    get: ->
      item = localStorage.getItem(key)
      JSON.parse(item) if item

$ ->
  window.lastLoad = now()
  window.refreshPassword = 30000
  window.tickerInterval = 200

  window.ticker = setInterval tick, window.tickerInterval
  $('#focusedInput').bind 'click', stopRefresh
  $('#focusedInput').bind 'blur', restartRefresh
  $('#optionSave').bind 'click', saveOptions

  loadOptions()
  loadPassword()

now = () ->
  d = new Date()
  d.getTime()

stopRefresh = () ->
  clearInterval(window.ticker)
  $('#focusedInput').select()
  false

restartRefresh = () ->
  window.lastLoad = now()
  window.ticker = setInterval tick, window.tickerInterval

setProgress = (perc) ->
  $('.progress-bar').css('width', "#{perc}%")

loadPassword = () ->
  options = loadOptions()
  $.get "/v1/getPassword?length=#{options.passwordLength}&special=#{options.useSpecial}", (data) ->
    $('#focusedInput').val(data)
    window.lastLoad = now()

saveOptions = () ->
  options =
    passwordLength: $('#passwordLengthOption').val()
    useSpecial: $('#useSpecialOption')[0].checked

  window.$storage('SecurePasswordOptions').set(options)
  $('#settingsModal').modal('hide')
  loadPassword()

loadOptions = () ->
  options = window.$storage('SecurePasswordOptions').get()
  if options == undefined
    options =
      passwordLength: 20
      useSpecial: false
  $('#passwordLengthOption').val(options.passwordLength)
  $('#useSpecialOption')[0].checked = options.useSpecial
  options

tick = () ->
  diff = now() - window.lastLoad
  perc = (window.refreshPassword - diff) / window.refreshPassword * 100
  setProgress perc
  if diff >= window.refreshPassword
    loadPassword()
