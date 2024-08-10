import { getSubjects } from './resources.js'

export async function showModal() {
  return $('body').append(`
    <div class="form-wrapper">
      <form id="appointment-form">
        <div class="form-close-button">
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
            <select id="form-subject" class="form-input" required>
              ${getSubjectElements()}
            </select>
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
              <label for="form-desc">Description<span class="required-mark">*</span></label>
              <textarea id="form-desc" class="form-input" required minlength="10" maxlength="1000"></textarea>
            </div>
          </div>
        </div>
        <div class="form-panel-io">
          <div class="error-box">
            <div class="error-icon">
              <svg class="icon-26jF1x" aria-hidden="true" role="img" xmlns="http://www.w3.org/2000/svg" width="3vh" height="3vh"
                fill="none" viewBox="0 0 24 24">
                <circle cx="12" cy="12" r="10" fill="transparent"></circle>
                <path fill="currentColor" fill-rule="evenodd"
                  d="M12 23a11 11 0 1 0 0-22 11 11 0 0 0 0 22Zm4.7-15.7a1 1 0 0 0-1.4 0L12 10.58l-3.3-3.3a1 1 0 0 0-1.4 1.42L10.58 12l-3.3 3.3a1 1 0 1 0 1.42 1.4L12 13.42l3.3 3.3a1 1 0 0 0 1.4-1.42L13.42 12l3.3-3.3a1 1 0 0 0 0-1.4Z"
                  clip-rule="evenodd" class=""></path>
              </svg>
            </div>
            <div>
              <span id="api-error"></span>
            </div>
          </div>
          <div class="buttons-row">
            <button id="modal-reset" type="reset">Cancel</button>
            <button id="modal-create" type="submit">Salvar</button>
          </div>
        </div>
      </form>
    </div>`)
}

async function getSubjectElements() {
  const subjs = await getSubjects()
  let list = ''

  for (const sub of subjs) {
    $(`<option value="${sub.id}">${sub.name}</option>`)
  }
  return list
}