import { ROUTES } from "../http/routes.js"
import CACHE from './cache.js'

let professors
let subjects

export async function updateCache() {
  const profs = (await get(ROUTES.GET_PROFESSORS)).professors
  const subs = (await get(ROUTES.GET_SUBJECTS)).subjects

  professors = profs.sort((a, b) => a.name.localeCompare(b.name))
  subjects = subs.sort((a, b) => a.name.localeCompare(b.name))
}

export function getProfessors() {
  return professors
}

export function getProfessorById(id) {
  for (const p of getProfessors())
    if (p['id'] === id)
      return p
  return null
}

export function getSubjects() {
  return subjects
}

export function getSubjectById(id) {
  for (const p of getSubjects()) {
    if (p['id'] === id) {
      return p
    }
  }
  return null
}

export async function fetchSelfUser() {
  return await get(ROUTES.GET_SELF_USER, getIdToken())
}

export async function fetchAppointments() {
  const resp = await get(ROUTES.GET_APPOINTMENTS, getIdToken())
  return resp == null ? {} : resp
}

export function getIdToken() {
  return localStorage.getItem(CACHE.ID_TOKEN)
}

export function getAccessToken() {
  return localStorage.getItem(CACHE.ACCESS_TOKEN)
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