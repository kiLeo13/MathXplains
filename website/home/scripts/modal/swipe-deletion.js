import { deleteAppointment } from "./delete-appointment.js"

export function registerSwipeDeletion() {
  $('.appointment').each((_, e) => {
    const appt = $(e)
    let startX, currentX, isSwiping = false // Fix positioning on Desktop
    
    if (!isDeletable(appt)) return

    appt.on('touchstart', (e) => {
      startX = e.originalEvent.touches[0].pageX
      isSwiping = true
    })

    appt.on('touchmove', (e) => {
      if (!isSwiping) return

      currentX = e.originalEvent.touches[0].pageX
      let diffX = currentX - startX

      if (diffX > 0) {
        appt.css('transform', `translateX(${diffX}px)`)
      }
    })

    appt.on('touchend', () => {
      if (!isSwiping) return

      isSwiping = false
      let diffX = currentX - startX

      if (diffX > 100) {
        appt.addClass('deleting')

        setTimeout(() => {
          deleteAppointment()
        }, 300)

      } else {
        appt.css('transform', 'translateX(0)')
      }
    })
  })
}

export function isDeletable(el) {
  const current = Date.now()
  const creation = new Date(el.attr('timestamp'))
  const hoursPassed = (current - creation) / (1000 * 60 * 60)

  return hoursPassed < 24
}