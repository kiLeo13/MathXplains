import { load, unload } from "../board/board.js"
import ROUTES from "../http/routes.js"
import { toggle } from "./buttons.js"

/**
 * This function is responsible for fetching the most up-to-date
 * information about notes from the server and automatically displaying
 * it on screen.
 */
export async function loadNotes() {
  const profile = sessionStorage.getItem('profile')
  const resp = await fetchNotes(profile)

  if (resp.ok) {
    const els = resp.notes
      .sort((a, b) => a.name.localeCompare(b.name))
      .map(buildNoteElement)

    displayNotes(els)
  } else {
    alert("Error: " + resp.message)
  }
}

export async function createNote(data) {
  const profile = sessionStorage.getItem('profile')
  const res = await postNote(profile, data.name)

  if (res.ok) {
    return res.note
  } else {
    alert('Não foi possível criar arquivo: ' + res.message)
    return null
  }
}

/**
 * This function sends an alert by blinking the profile textarea,
 * meaning the user must feed this field with some data.
 */
export function requireProfile() {
  const $profile = $('#profile-in')
  $profile.trigger('focus')

  $profile.css({
    "box-shadow": "0 0 30px rgb(200, 107, 107)",
    "border-color": "rgb(177, 107, 107)"
  })
  
  setTimeout(() => {
    $profile.css({
      "box-shadow": "none",
      "border-color": "gray"
    })
  }, 1000)
}

async function postNote(profile, name) {
  const resp = await ROUTES.CREATE_NOTE.send({
    body: {
      name: name,
      profile: profile
    }
  })
  const json = await resp.json()

  return {
    ok: resp.ok,
    note: json,
    message: json.message // Only available on failures
  }
}

async function fetchNotes(profile) {
  const resp = await ROUTES.LIST_NOTES.send({
    query: {profile: profile}
  })
  const json = await resp.json()
  
  return {
    ok: resp.ok,
    notes: json.notes,
    message: json.message // Only available on failures
  }
}

// It is expected "els" to be raw strings, not
// real elements built with jQuery.

function displayNotes(els) {
  const $container = $('.sidebar-notes')

  $container.empty()

  for (const el of els) {
    const $note = $(el)

    setListeners($note)
    $container.append($note)
  }
}

function setListeners($el) {
  $el.find('.delete-file-button').on('click', deleteNote)
  $el.on('click', loadNoteHandler)
}

async function loadNoteHandler(e) {
  const $item = $(e.currentTarget)
  const id = $item.attr('id').substring('note-'.length)
  const resp = await openNote(id)

  if (resp.status === 200) {
    load(id, resp.body.name, resp.body.content)
    return
  }

  if (resp.status === 404) {
    $el.remove()
  }
}

async function deleteNote(e) {
  const $btn = $(e.currentTarget)
  const profile = sessionStorage.getItem('profile')
  const id = $btn.attr('id').substring('delete-item-'.length)

  toggle($btn, false)
  const resp = await ROUTES.DELETE_NOTE.send({
    path: {id: id},
    headers: {Profile: profile}
  })
  toggle($btn, true)
  
  // If its not found, then the resource was previously deleted,
  // we can just proceed to remove it from the UI
  if (resp.ok || resp.status === 404) {
    $btn.closest('.sidebar-file-item').remove()

    if (isCurrent(id)) unload(id)
  } else {
    const json = await resp.json()
    alert("Error: " + json.message)
  }
}

function isCurrent(id) {
  const $board = $('#board-content')
  const boardId = $board.attr('note')
  
  return boardId === id
}

async function openNote(id) {
  const resp = await ROUTES.OPEN_NOTE.send({
    path: {id: id},
    query: {profile: sessionStorage.getItem('profile')}
  })

  return {
    status: resp.status,
    body: await resp.json()
  }
}

function buildNoteElement(data) {
  return `
    <div class="sidebar-file-item" id="note-${data.id}">
      <span class="sidebar-file-title">${data.name}</span>
      <div class="file-button-wrapper">
        <button class="delete-file-button" id="delete-item-${data.id}">
          <svg version="1.0" xmlns="http://www.w3.org/2000/svg" width="20px" height="20px" viewBox="0 0 512.000000 512.000000"
            preserveAspectRatio="xMidYMid meet">
            <g transform="translate(0.000000,512.000000) scale(0.100000,-0.100000)" stroke="none">
              <path
                d="M1835 4671 c-168 -78 -164 -314 7 -385 33 -14 125 -16 720 -16 643 0 685 2 723 19 165 76 165 306 0 382 -38 17 -79 19 -725 19 -646 0 -687 -2 -725 -19z" />
              <path
                d="M798 3830 c-143 -43 -203 -221 -112 -329 58 -70 75 -76 234 -81 l145 -5 5 -1140 c3 -627 9 -1151 14 -1165 4 -14 17 -54 28 -90 96 -302 361 -530 678 -580 88 -13 1452 -13 1540 0 300 48 567 265 665 540 53 147 49 68 55 1295 l5 1140 145 5 c159 5 176 11 234 81 77 91 47 243 -62 306 l-47 28 -1750 2 c-962 1 -1762 -2 -1777 -7z m1449 -876 c34 -21 55 -44 75 -83 l28 -53 0 -675 c0 -477 -3 -687 -11 -714 -56 -184 -318 -200 -400 -24 -17 38 -19 80 -19 729 0 682 0 690 21 733 56 115 196 155 306 87z m804 25 c52 -16 103 -61 128 -112 21 -43 21 -54 21 -733 0 -649 -2 -691 -19 -729 -83 -178 -349 -159 -400 29 -9 29 -11 235 -9 722 l3 681 30 48 c22 36 45 57 84 77 57 30 102 35 162 17z" />
            </g>
          </svg>
        </button>
      </div>
    </div>`
}