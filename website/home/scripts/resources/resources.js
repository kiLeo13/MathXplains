const PROFESSORS_URL   = `${window.origin}/api/professors?known=true`
const SUBJECTS_URL     = `${window.origin}/api/subjects?available=true`
const APPOINTMENTS_URL = `${window.origin}/api/appointments`
const SELF_USER_URL    = `${window.origin}/api/users/@me`

let professors
let subjects

export async function getProfessors() {
  return !professors
    ? (await get(PROFESSORS_URL)).professors
    : professors
}

export async function getSubjects() {
  return !subjects
    ? (await get(SUBJECTS_URL)).subjects
    : subjects
}

export async function fetchSelfUser() {
  return (await get(SELF_USER_URL, getCacheToken()))
}

export async function fetchAppointments() {
  return {count: 0, appointments: [], max: 100} //(await get(APPOINTMENTS_URL, getCacheToken()))
}

export function getCacheToken() {
  return localStorage.getItem('authToken')
}

async function get(url, auth) {
  const resp = await fetch(url, {
    method: 'GET',
    headers: !auth ? {} : getHeader(auth)
  })

  return await resp.json()
}

function getHeader(auth) {
  return {
    Authorization: `Bearer ${auth}`
  }
}