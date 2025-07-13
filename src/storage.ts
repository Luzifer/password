const defaultSettings = {
  passwordLength: 32,
  useSpecial: false,
  useXKCD: false,
  xkcdSeparator: '',
}

const storageKey = 'SecurePasswordOptions'

export default {
  set: (value: any): void => {
    localStorage.setItem(storageKey, JSON.stringify(value))
  },
  get: (): any => {
    let item = localStorage.getItem(storageKey)
    if (!item) { return defaultSettings }
    return JSON.parse(item)
  }
}
