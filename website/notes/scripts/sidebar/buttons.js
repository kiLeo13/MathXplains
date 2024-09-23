export function enableCreate() {
  toggle($('#new-file-button'), true)
}

export function disableCreate() {
  toggle($('#new-file-button'), false)
}

export function enableReload() {
  toggle($('#refresh-files-button'), true)
}

export function disableReload() {
  toggle($('#refresh-files-button'), false)
}

/**
 * This function does not reload any elements, it simply shows the loading
 * icon on the refresh button.
 * 
 * Keep in mind that this function ALSO DISABLES THE BUTTON.
 * 
 * @param {*} flag whether it should display or hide the loading icon. Defaults to `true`.
 */
export function loadingRefresh(flag = true) {
  const $btn = $('#new-file-button')
  setLoading($btn, flag)
}

/**
 * This function does not reload any elements, it simply shows the loading
 * icon on the create button.
 * 
 * Keep in mind that this function ALSO DISABLES THE BUTTON.
 * 
 * @param {*} flag whether it should display or hide the loading icon. Defaults to `true`.
 */
export function loadingCreate(flag = true) {
  const $btn = $('#refresh-files-button')
  setLoading($btn, flag)
}

export function toggle($btn, flag) {
  if (flag) {
    $btn.removeAttr('disabled')
  } else {
    $btn.attr('disabled', 'true')
  }
}

function setLoading($btn, flag) {
  toggle($btn, !flag)
  const $svg = $btn.prev('.loader')

  if (flag) {
    $svg.show()
  } else {    
    $svg.hide()
  }
}