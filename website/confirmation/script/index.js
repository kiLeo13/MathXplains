import { validateSubmitButton } from "./resources/buttons.js"

$(() => {
  const url = new URL(window.location.href)
  const email = url.searchParams.get('e')
  const code = url.searchParams.get('c')
  
  if (email) {
    $('#email').val(email)
    url.searchParams.delete('e')
  }

  if (code) {
    $('#code').val(code)
    url.searchParams.delete('c')
  }

  if (code && email) {
    validateSubmitButton()
    $('#submit-button').trigger('click')
  }

  window.history.replaceState({}, '', url.href)
})