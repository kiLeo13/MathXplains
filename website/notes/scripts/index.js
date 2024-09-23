import { disableCreate, disableReload } from "./sidebar/buttons.js"

const WELCOME_CACHE_KEY = 'welcome-screen'

$(() => {
  ensureWelcomeScreen()
  disableReload()
  disableCreate()
})

function ensureWelcomeScreen() {
  if (hasReadWelcome()) return
  
  const screen = buildWelcomeScreen()
  $('body').append(screen)
  profileFocus(false)
  
  setTimeout(loadTitleShadows, 1700)
  setListeners()
}

function loadTitleShadows() {
  $('.domain-title').css('box-shadow', '0 0 20px rgba(255, 87, 51, 0.3)')
  $('.app-title').css('box-shadow', '0 0 20px rgba(143, 255, 51, 0.3)')
}

function setListeners() {
  const $wrapper = $('.welcome-wrapper')
  const closeWelcome = () => {
    $('.welcome-screen').css('scale', '0')

    setTimeout(() => {
      $wrapper.remove()
      setWelcomeRead()
      profileFocus()
    }, 300)
  }

  $('.welcome-close-action').on('click', closeWelcome)
  $('body').on('keydown', (e) => {
    if (e.key === 'Escape') closeWelcome()
  })
}

function setWelcomeRead() {
  localStorage.setItem(WELCOME_CACHE_KEY, 'READ')
}

function hasReadWelcome() {
  return (localStorage.getItem(WELCOME_CACHE_KEY) || '').toUpperCase() === 'READ'
}

function profileFocus(flag = true) {
  const $profile = $('#profile-in')

  if (flag) {
    $profile.trigger('focus')
  } else {
    $profile.trigger('blur')
  }
}

function buildWelcomeScreen() {
  return `
  <div class="welcome-wrapper">
    <div class="welcome-screen">
      <div class="welcome-body">
        <header class="welcome-head">
          <h1 class="welcome-title">
            <span class="domain-title">MathXplains</span>
            <span class="app-title">Notes</span>
          </h1>
          <span class="welcome-subtitle">Bem-vindo&lpar;a&rpar; à área de anotações MathXplains!</span>
          <div class="welcome-division"></div>
        </header>
        <main class="welcome-main">
          <div class="category">
            <h2>Geral</h2>
            <p class="category-p">
              Considere este projeto um bloco de notas compartilhado e, talvez,
              aprimorado, não espere uma alternativa para <a href="https://www.microsoft.com/microsoft-365/word" target="_blank" rel="noopener noreferrer">Microsoft Word</a>.
            </p>
            <p class="category-p">
              Mesmo não tendo um limite para quantos arquivos você pode criar, há um limite de <code>500.000</code> caracteres por arquivo.
            </p>
            <p class="category-p">
              Sim, a interface foi inspirada no ChatGPT.
            </p>
          </div>
          <div class="category">
            <h2>Segurança</h2>
            <p class="category-p">
              Todos os perfis utilizados neste site são armazenados criptograficamente,
              utilizando hashing <a href="https://en.wikipedia.org/wiki/SHA-2" target="_blank" rel="noopener noreferrer">SHA-256</a>.
              <br>
              <span class="disclaimer">Nem eu tenho acesso às senhas</span>
            </p>
          </div>
        </main>
      </div>

      <div class="close-welcome-x welcome-close-action">
        <svg version="1.0" xmlns="http://www.w3.org/2000/svg" width="12px" height="12px"
          viewBox="0 0 512.000000 512.000000" preserveAspectRatio="xMidYMid meet">
      
          <g transform="translate(0.000000,512.000000) scale(0.100000,-0.100000)" stroke="none">
            <path d="M271 5109 c-104 -20 -194 -91 -239 -187 -22 -47 -27 -71 -27 -137 0
      -155 -68 -78 1075 -1220 l1010 -1010 -1019 -1020 c-827 -829 -1022 -1029
      -1041 -1070 -34 -71 -35 -199 -2 -270 29 -63 93 -129 157 -163 72 -37 187 -42
      270 -12 58 22 94 56 1083 1044 l1022 1021 1018 -1016 c817 -816 1027 -1021
      1067 -1040 42 -20 65 -24 145 -24 83 0 101 4 145 27 62 32 129 103 158 166 32
      68 30 197 -3 267 -19 41 -214 241 -1041 1070 l-1019 1020 1010 1010 c1143
      1142 1075 1065 1075 1220 0 65 -5 90 -26 135 -81 172 -284 242 -454 158 -26
      -13 -391 -370 -1057 -1036 l-1018 -1017 -1017 1017 c-667 666 -1032 1023
      -1058 1036 -34 17 -145 45 -164 41 -3 -1 -26 -5 -50 -10z" />
          </g>
        </svg>
      </div>
      <div class="close-welcome-ok welcome-close-action">Entendi</div>
    </div>
  </div>`
}