/**
 * Formats a number of seconds into a string of minutes and seconds.
 * @param {number} secs - Total number of seconds.
 * @returns {string} Time formatted as "M:SS".
 */
function formatTime(secs) {
    const minutes = Math.floor(secs / 60)
    const seconds = secs % 60
    const sencondsPadded = seconds < 10 ? "0" + seconds : seconds

    return `${minutes} : ${sencondsPadded}`
}

/**
 * Gets the current grid size from the URL search params.
 * @returns {string} Grid size.
 */
function getSize() {
    const params = new URLSearchParams(location.search)
    const size = params.get("size")
    return size
}

/**
 * Stops the timer and saves it to localStorage.
 * @param {number} interval - Timer interval (returned from `setInterval`).
 * @param {number} secs - Number of elapsed seconds to be saved.
 */
function stopAndSave(interval, size, secs) {
    clearInterval(interval)
    let times = JSON.parse(localStorage.getItem("times"))

    if (!times) {
        times = {}
    }

    if (!times[size]) {
        times[size] = []
    }

    times[size].push(secs)
    localStorage.setItem("times", JSON.stringify(times))
}

/**
 * Total number of seconds elapsed.
 */
let secs = 0


/**
 * The timer HTML element, or `null` if it doesn't exist.
 */
const timerElement = document.getElementById("timer");


/**
 * The interval that updates elapsed seconds, and timer element if exists.
 */
const intv = setInterval(() => {
    secs++
    if (timerElement) {
        timerElement.textContent = formatTime(secs)
    }
}, 1000);


document.getElementById("done").addEventListener("click", () => {
    stopAndSave(intv, getSize(), secs)
})
