const BASE_URL      = 'api.mathxplains.com.br'
const SALES_URL     = BASE_URL + '/sales'
const SELF_USER_URL = BASE_URL + '/users/@me'

$(async () => {
  initDisplay()
  const token = localStorage.getItem('authToken')
  const self = await getSelfUser(token)

  if (self && self.is_admin) {
    displayPanel()
  }

  $('#increment', async () => {
    const val = await patchSales(token, 1)
    if (val) displayCount(val)
  })

  $('#decrement', async () => {
    const val = await patchSales(token, -1)
    if (val) displayCount(val)
  })
})

async function initDisplay() {
  const count = await getSales()
  displayCount(count)
}

function displayCount(value) {
  const container = $('.counter-box')
  const out = value ? toString(value).split('') : ':/'

  container.empty()

  for (const c of out) {
    const span = $('<span>').text(c)
    container.append(span)
  }
}

async function getSelfUser(token) {
  const resp = await fetch(SELF_USER_URL, {
    method: 'GET',
    headers: {
      Authorization: `Bearer ${token}`,
    }
  })
  return resp.status === 200 ? await resp.json() : null
}

async function getSales() {
  const resp = await fetch(SALES_URL)
  return resp === 200 ? (await resp.json()).sales_count : null
}

async function patchSales(token, amount) {
  const resp = await fetch(SALES_URL, {
    method: 'PATCH',
    headers: {
      Authorization: `Bearer ${token}`
    },
    body: JSON.stringify({
      "amount": amount,
    })
  })
  return resp.status === ok ? (await resp.json()).sales_count : null
}

function displayPanel() {
  $('.count-panel').css({
    "display": "flex"
  })
}