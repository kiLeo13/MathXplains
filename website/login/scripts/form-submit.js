import { setButtonLoading, isDisabled } from "./buttons.js"

$(() => {
  const form = $('#login-form')

  form.on('submit', (e) => {
    e.preventDefault()

    if (isDisabled()) return

    const email = $('#email').val()
    const password = $('#password').val()

    setButtonLoading(true)

    handleRequest({
      email: email,
      password: password
    })
  })
})

async function handleRequest(credentials) {
  const resp = await fetch(`${location.origin}/api/users/login`, {
    method: 'POST',
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      email: credentials.email,
      password: credentials.password
    })
  })
  const json = await resp.json()

  if (json.token) {
    localStorage.setItem('authToken', json.token)
    location.href = `${location.origin}/home`
  } else {
    setButtonLoading(false)
    displayError(json.error)
  }
}

function displayError(msg) {
  $('#error-message').css({
    opacity: 1
  }).text(msg)
}