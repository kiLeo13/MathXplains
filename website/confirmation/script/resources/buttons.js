export function setButtonLoading(flag) {
  setSubmitButtonClickable(!flag)

  $('#loader-icon').css({
    "opacity": flag ? 1 : 0
  })
}

export function validateSubmitButton() {
  const inputs = $('.form-input')
  let valid = true

  inputs.each((_, e) => {
    const element = $(e)

    if (!element.is(':valid')) {
      valid = false
      return
    }
  })

  setSubmitButtonClickable(valid)
}

export function canSubmit() {
  return isEnabled('#submit-button')
}

export function canResend() {
  return isEnabled('#resend-code')
}

export function showError(msg) {
  $('#error-message').css({
    opacity: 1
  }).text(msg)
}

export function hideError() {
  $('#error-message')
    .css({"opacity": 0})
    .val('')
}

export function setSubmitButtonClickable(flag) {
  setButtonClickable('#submit-button', flag)
}

export function setResendButtonClickable(flag) {
  setButtonClickable('#resend-code', flag)
}

export function sendResendCooldown(millis = 30000) {
  setResendButtonClickable(false)
  setTimeout(() => {
    setResendButtonClickable(true)
  }, millis)
}

function isEnabled(name) {
  return !$(name).hasClass("disabled")
}

function setButtonClickable(button, flag) {
  const el = $(button)

  if (flag) el.removeClass('disabled')
  else el.addClass('disabled')
}