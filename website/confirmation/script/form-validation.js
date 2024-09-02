import { sendResendCooldown, setSubmitButtonClickable, validateSubmitButton } from './resources/buttons.js'

$(() => {
  setSubmitButtonClickable(false)
  sendResendCooldown(5000)
  $('.form-input').on('keyup', validateSubmitButton)
})