import { getSubjects } from './resources.js'
import { onSubmit } from '../modal/submit-appointment.js'

/**
 * Removes the modal completely from screen.
 * If the modal is not being displayed, then this is a no-op function.
 */
export function closeModal() {
  const form = $('#form-wrapper')
  if (form) form.remove()
}

/**
 * Displays the modal on screen.
 * If there is already a modal being displayed, it will be removed.
 * 
 * Subjects in the <select> field will be lazily-loaded.
 */
export async function showModal() {
  closeModal()
  const form = `
    <div id="form-wrapper">
      <form id="appointment-form">
        <div class="close-modal-icon-container">
          <svg class="closeIcon_f9a4c9" aria-hidden="true" role="img" xmlns="http://www.w3.org/2000/svg" width="3vh" height="3vh"
            fill="none" viewBox="0 0 24 24">
            <path fill="currentColor"
              d="M17.3 18.7a1 1 0 0 0 1.4-1.4L13.42 12l5.3-5.3a1 1 0 0 0-1.42-1.4L12 10.58l-5.3-5.3a1 1 0 0 0-1.4 1.42L10.58 12l-5.3 5.3a1 1 0 1 0 1.42 1.4L12 13.42l5.3 5.3Z">
            </path>
          </svg>
        </div>
        <div class="form-contents">
          <div class="form-field">
            <label for="form-subject">Matéria<span class="required-mark">*</span></label>
            <select id="form-subject" class="form-input" required>${buildSubjects()}</select>
          </div>
          <div class="form-field">
            <label for="form-topic">Assunto<span class="required-mark">*</span></label>
            <input type="text" id="form-topic" class="form-input" required minlength="5" maxlength="30">
          </div>
          <div class="form-field">
            <label for="form-date">Data<span class="required-mark">*</span></label>
            <input type="date" id="form-date" class="form-input" required>
          </div>
          <div class="description-container">
            <div class="form-field">
              <label for="form-desc">Descrição<span class="required-mark">*</span></label>
              <textarea id="form-desc" class="form-input" required minlength="10" maxlength="1000"></textarea>
            </div>
          </div>
        </div>
        <div id="form-panel-io">
          <div class="buttons-row">
            <button id="modal-reset" type="reset">Limpar</button>
            <div class="modal-submit-container">
              <button id="modal-create" type="submit">Salvar</button>
              <span class="loader-icon"></span>
            </div>
          </div>
        </div>
      </form>
    </div>`
  
  $('body').append(form)
  createModalListeners()
}

function buildSubjects() {
  let html = ''

  for (const sub of getSubjects()) {
    html += `<option value="${sub.id}" class="subject-option">${sub.name}</option>`
  }
  return html
}

/**
 * Coalesces an error message inside the modal.
 * If no modals are being displayed, this is a no-op function.
 * 
 * If there is already an error being displayed, its content will be overwritten.
 */
export function showError(msg) {
  const err = `
    <div id="modal-error-container" class="error-box">
      <div class="error-icon">
        <svg class="icon-26jF1x" aria-hidden="true" role="img" xmlns="http://www.w3.org/2000/svg" width="3vh" height="3vh"
          fill="none" viewBox="0 0 24 24">
          <circle cx="12" cy="12" r="10" fill="transparent"></circle>
          <path fill="currentColor" fill-rule="evenodd"
            d="M12 23a11 11 0 1 0 0-22 11 11 0 0 0 0 22Zm4.7-15.7a1 1 0 0 0-1.4 0L12 10.58l-3.3-3.3a1 1 0 0 0-1.4 1.42L10.58 12l-3.3 3.3a1 1 0 1 0 1.42 1.4L12 13.42l3.3 3.3a1 1 0 0 0 1.4-1.42L13.42 12l3.3-3.3a1 1 0 0 0 0-1.4Z"
            clip-rule="evenodd"></path>
        </svg>
      </div>
      <div>
        <span id="api-error">${msg}</span>
      </div>
    </div>`

  const panel = $('#form-panel-io')

  removeError()
  if (panel) panel.prepend(err)
  
  fitError()
}

/**
 * This function does not resize the error box itself (#modal-error-container),
 * instead, it adds bottom padding to the div containing the inputs, in order to fit the entire error.
 */
function fitError() {
  const errorHeight = $('#modal-error-container').height()
  const form = $('#appointment-form')

  form.css({"padding-bottom": `calc(${errorHeight}px + 10vh`})
}

/**
 * Removes the error from the modal
 */
export function removeError() {
  $('#modal-error-container').remove()
}

export function setSaveLoading(flag = true) {
  const but = $('#modal-create')

  if (flag) {
    but.prop("disabled", true)
  } else {
    but.removeProp("disabled")
  }
}

function createModalListeners() {
  $(document).on('keydown', (e) => {
    if (e.key === 'Escape') closeModal()
  })
  
  $('.close-modal-icon-container').on('click', () => {
    closeModal()
  })

  $('#appointment-form').on('submit', onSubmit)
}