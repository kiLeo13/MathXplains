import { setButtonClickable } from './buttons.js'

$(() => {
  setButtonClickable(false)
  $('.form-input').on('keyup', buttonValidationHandler)
})

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