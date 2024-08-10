export function setButtonLoading(flag) {
  setButtonClickable(!flag)

  $('#loader-icon').css({
    "opacity": flag ? 1 : 0
  })
}

export function setButtonClickable(flag) {
  const button = $('#submit-button')

  button.css({
    "opacity": flag ? 1 : 0.3,
    "cursor": flag ? "pointer" : "not-allowed",
  })
}

export function isDisabled() {
  return $('#submit-button').css('cursor') === 'not-allowed'
}