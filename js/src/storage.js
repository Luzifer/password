export default {
  set: (key, value) => {
    localStorage.setItem(key, JSON.stringify(value))
  },
  get: (key) => {
    let item = localStorage.getItem(key)
    if (item) return JSON.parse(item)
  }
}
