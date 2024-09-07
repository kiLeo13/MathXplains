import { signOut } from "../http/sign-out.js"

$(() => {
  const $container = $('.user-container')

  $container.on('click', () => {
    const $userActions = $('.user-actions')

    if ($userActions.length) {
      $userActions.remove()
    } else {
      $container.append(getHTMLUserTools())
      registerUserToolsListeners()
    }
  })
})

function registerUserToolsListeners() {
  $('.sign-out-actions').on('click', signOut)
}

function getHTMLUserTools() {
  return `
    <div class="user-actions">
      <div class="sign-out-actions">
        <svg height="32px" width="32px" version="1.1" id="sign-out-icon" xmlns="http://www.w3.org/2000/svg"
          xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="-23.1 -23.1 431.17 431.17" xml:space="preserve"
          stroke-width="23.09826">
          <g id="SVGRepo_bgCarrier" stroke-width="0" />
          <g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"
            stroke-width="1.539884" />
          <g id="SVGRepo_iconCarrier">
            <g>
              <g id="Sign_Out">
                <path
                  d="M180.455,360.91H24.061V24.061h156.394c6.641,0,12.03-5.39,12.03-12.03s-5.39-12.03-12.03-12.03H12.03 C5.39,0.001,0,5.39,0,12.031V372.94c0,6.641,5.39,12.03,12.03,12.03h168.424c6.641,0,12.03-5.39,12.03-12.03 C192.485,366.299,187.095,360.91,180.455,360.91z" />
                <path
                  d="M381.481,184.088l-83.009-84.2c-4.704-4.752-12.319-4.74-17.011,0c-4.704,4.74-4.704,12.439,0,17.179l62.558,63.46H96.279 c-6.641,0-12.03,5.438-12.03,12.151c0,6.713,5.39,12.151,12.03,12.151h247.74l-62.558,63.46c-4.704,4.752-4.704,12.439,0,17.179 c4.704,4.752,12.319,4.752,17.011,0l82.997-84.2C386.113,196.588,386.161,188.756,381.481,184.088z" />
              </g>
            </g>
          </g>
        </svg>
        <span id="sign-out-button">Deslogar</span>
      </div>
    </div>`
}