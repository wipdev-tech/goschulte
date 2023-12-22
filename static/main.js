/**
 * Formats a number of seconds into a string of minutes and seconds.
 * @param {number} secs - The total number of seconds.
 * @returns {string} A string in the format "M:SS".
 */
function formatTime(secs) {
    const minutes = Math.floor(secs / 60)
    const seconds = secs % 60
    const sencondsPadded = seconds < 10 ? "0" + seconds : seconds

    return `${minutes} : ${sencondsPadded}`
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
    clearInterval(intv)
})
