import ROUTES from "../http/routes.js"

export function addFile() {

}

export async function loadFiles(profile) {

}

export async function createNote(data) {
  const profile = sessionStorage.getItem('profile')
  const res = await postNote(profile, data.name, data.content)

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

async function postNote(profile, name, content = '') {
  const resp = await fetch(ROUTES.CREATE_NOTE, {
    method: "POST",
    body: JSON.stringify({
      name: name,
      profile: profile,
      content: content
    }),
    headers: {
      "Content-Type": "application/json"
    }
  })
  const json = await resp.json()

  return {
    ok: resp.ok,
    note: json,
    message: json.message // Only available on failures
  }
}

async function getNotes(profile) {
  const resp = await fetch(ROUTES.GET_NOTES + profile)
  const json = await resp.json()
  
  return {
    ok: resp.ok,
    notes: json.notes,
    message: json.message // Only available on failures
  }
}