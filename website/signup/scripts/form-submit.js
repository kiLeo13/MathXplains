$(() => {
  const form = $('form')
  setButtonClickable(false)
  $('.form-input').on('keyup', buttonValidationHandler)

  form.on('submit', async (e) => {
    e.preventDefault()

    if (isDisabled()) return

    const name = $('#name').val()
    const email = $('#email').val()
    const password = $('#password').val()

    setButtonLoading(true)

    setTimeout(() => {
      setButtonLoading(false)
    }, 3000)

    const resp = await requestCreation({
      name: name,
      email: email,
      password: password,
    })

    if (resp.error) {
      displayError(resp.error)
    } else {
      
    }
  })
})

function isDisabled() {
  return $('#submit-button').css('cursor') === 'not-allowed'
}

async function requestCreation(body) {
  const resp = await fetch(`${location.origin}/api/users`, {
    method: 'POST',
    body: body,
    headers: { "Content-Type": "application/json" }
  })

  return await resp.json()
}

function displayError(msg) {
  $('#error-message').css({
    "opacity": 1
  }).val(msg)
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