/**
 * @typedef {number[]} numbers
 */

/**
 * Calculates the mean of an array of numbers.
 * @param {numbers} ar - an array of numbers
 * @returns {number} Calculated mean.
 */
function mean(ar) {
    const sum = ar.reduce((acc, e) => acc + e)
    return sum / ar.length
}

const times = JSON.parse(localStorage.getItem("times"))

for (size in times) {
    const m = mean(times[size])
    const liNew = document.createElement("li")
    liNew.classList.add("text-lg")
    liNew.innerHTML = `<b>${size}x${size}</b>: ${m}s`
    document.getElementById("stats").append(liNew)
}
