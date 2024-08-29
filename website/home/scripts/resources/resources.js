import ROUTES from "../http/routes.js"

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
  // return {
  //   "active": 20,
  //   "appointments": [
  //     {
  //       "id": 1,
  //       "topic": "Prova de AFO",
  //       "description": "Gostaria de ajuda na prova de AFO, que eu com certeza vou tirar outro I.",
  //       "user_id": "515b25a0-d021-70c2-5693-60861181f48c",
  //       "subject_id": 11,
  //       "professor_id": 16,
  //       "rejected": false,
  //       "is_active": false,
  //       "scheduled_at": "2024-08-24",
  //       "created_at": "2024-08-25T23:08:51Z",
  //       "updated_at": "2024-08-25T23:08:51Z"
  //     },
  //     {
  //       "id": 3,
  //       "topic": "Probabilidade",
  //       "description": "Matheus que gosta dessa matéria, então vou aproveitar e pedir uma explicação para ele sobre isso, só espero que não caia isso na prova, amém.",
  //       "user_id": "515b25a0-d021-70c2-5693-60861181f48c",
  //       "subject_id": 5,
  //       "professor_id": 113,
  //       "rejected": false,
  //       "is_active": true,
  //       "scheduled_at": "2024-08-29",
  //       "created_at": "2024-08-28T07:25:59Z",
  //       "updated_at": "2024-08-26T14:47:59Z"
  //     },
  //     {
  //       "id": 4,
  //       "topic": "Física Quântica",
  //       "description": "Espero de verdade que isso não caia no ENEM.",
  //       "user_id": "515b25a0-d021-70c2-5693-60861181f48c",
  //       "subject_id": 16,
  //       "professor_id": null,
  //       "rejected": false,
  //       "is_active": true,
  //       "scheduled_at": "2024-08-29",
  //       "created_at": "2024-08-28T19:25:59Z",
  //       "updated_at": "2024-08-26T14:54:12Z"
  //     }
  //   ],
  //   "max": 20
  // }
 return await get(ROUTES.GET_APPOINTMENTS, getIdToken())
}

export function getIdToken() {
  return localStorage.getItem('idToken')
}

export function getAccessToken() {
  return localStorage.getItem('accessToken')
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