import ROUTES from './http/routes.js'

$(() => {
  setButtonClickable(false)
  $('.form-input').on('keyup', buttonValidationHandler)

  $('form').on('submit', async (e) => {
    e.preventDefault()

    if (isDisabled()) return

    const name = $('#name').val()
    const email = $('#email').val()
    const password = $('#password').val()

    setButtonLoading(true)
    hideError()
    handleCreation({
      name: name,
      email: email,
      password: password,
    })
  })
})

function isDisabled() {
  return $('#submit-button').css('cursor') === 'not-allowed'
}

async function handleCreation(body) {
  const resp = await fetch(ROUTES.CREATE_USER, {
    method: "POST",
    body: JSON.stringify(body),
    headers: {
      "Content-Type": "application/json"
    }
  })
  
  if (resp.ok) {
    loadEmailConfirmationPage(body.email)
  } else {
    const json = await resp.json()

    displayError(json.message)
    setButtonLoading(false)
  }
}

async function loadEmailConfirmationPage(email) {
  const url = '../../confirmation/index.html?e=' + encodeURI(email)
  const resp = await fetch(url)

  console.log(JSON.stringify(resp))

  if (resp.ok) {
    $('body').load(url)
  
    history.pushState(null, null, url)
  } else {
    location.href = url
  }
}

function displayError(msg) {
  $('#error-message').css({
    "opacity": 1
  }).text(msg)
}

function hideError() {
  $('#error-message').css({
    "opacity": 9
  }).text('')
}

function buttonValidationHandler() {
  const inputs = $('.form-input')
  let valid = true

  inputs.each((_, e) => {

    const element = $(e)

    if (!element.is(':valid')) {
      valid = false
      return
    }
  })

  setButtonClickable(valid)
}

function setButtonClickable(flag) {
  const button = $('#submit-button')

  button.css({
    "opacity": flag ? 1 : 0.3,
    "cursor": flag ? "pointer" : "not-allowed",
  })
}

function setButtonLoading(flag) {
  setButtonClickable(!flag)

  $('#loader-icon').css({
    "opacity": flag ? 1 : 0
  })
}