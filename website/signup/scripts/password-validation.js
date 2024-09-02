const PASSWORD_MIN_LENGTH = 8
const PASSWORD_SPECIAL_CHARS_PATTERN = /[@$!%*?&\-_]/
const PASSWORD_NUMBER_PATTERN = /\d/

const RED   = 'rgba(255, 0, 0, 0.3)'
const GREEN = 'rgba(0, 255, 0, 0.3)'

$(() => {
  updatePasswordField()

  $('#password').on('keyup', () => {
    updatePasswordField()
    updateTips()
  })

  $('#password').on('focus focusout', (e) => {
    updateTips()
  })
})

function updatePasswordField() {
  const password = $("#password").val()

  const minChars     = $('#min-chars')
  const capitalChars = $('#capital-chars')
  const minNumbers   = $('#min-numbers')
  const specialChars = $('#special-chars')

  const checks = checkPassword(password)

  minChars.css({    "background-color": checks.length   ? GREEN : RED })
  capitalChars.css({"background-color": checks.cases    ? GREEN : RED })
  minNumbers.css({  "background-color": checks.number   ? GREEN : RED })
  specialChars.css({"background-color": checks.specials ? GREEN : RED })
}

function checkPassword(password) {
  return {
    length: checkPasswordLength(password),
    cases: checkPasswordCase(password),
    number: checkPasswordNumber(password),
    specials: checkPasswordSpecialChars(password)
  }
}

function checkPasswordLength(input) {
  return input.length >= PASSWORD_MIN_LENGTH
}

function checkPasswordCase(input) {
  return input.toLowerCase() !== input && input.toUpperCase() !== input
}

function checkPasswordNumber(input) {
  return PASSWORD_NUMBER_PATTERN.test(input)
}

function checkPasswordSpecialChars(password) {
  return PASSWORD_SPECIAL_CHARS_PATTERN.test(password)
}

function updateTips() {
  const passwordInput = $('#password')
  const input = $('#password-requirements')
  const displayable = passwordInput.is(':invalid') && passwordInput.is(':focus')

  input.css({ "transform": `scaleY(${displayable ? 100 : 0}%)` })
}