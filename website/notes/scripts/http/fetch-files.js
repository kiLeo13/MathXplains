import ROUTES from './routes.js'

export async function fetchFiles(profile) {
  const resp = await fetch(ROUTES.GET_NOTES + profile)
  const json = await resp.json()

  if (resp.ok) {
    return {
      ok: true,
      notes: json.notes
    }
  } else {
    return {
      ok: false,
      message: json.message
    }
  }
}