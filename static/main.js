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
 * The number of seconds to format.
 * @type {number}
 */
let secs = 0

setInterval(() => {
    secs++
    document.getElementById("timer").textContent = formatTime(secs)
}, 1000);

