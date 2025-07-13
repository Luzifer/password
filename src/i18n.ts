import { createI18n } from "vue-i18n";

/*
 * Adding a new language
 *
 * - Copy i18n/en.json into i18n/<lc>.json
 * - Translate the values of that file
 * - Add an import below and add the translation to the `messages` const
 *
 * Lanugages should either be added as their 2-letter-code (i.e. `de`)
 * or if there are specific variants for countries (i.e. `pt-BR` / `pt-PT`)
 * they are added as `pr-br` in the `messages` const.
 */

import en from './i18n/en.json'

const messages = {
  en,
}

function getUserLocale(): string {
  // Try URL search param
  const searchParams = new URLSearchParams(window.location.search)
  if (searchParams.has('lang')) {
    return searchParams.get('lang') as string
  }

  // Try cookie
  const cookieCollection = Object.fromEntries(document.cookie.split('; ').map(c => c.split('=')))
  if (cookieCollection.lang) {
    return cookieCollection.lang
  }

  // Navigator knows best, right?
  return navigator.language
}

const i18n = createI18n({
  locale: getUserLocale(),
  fallbackLocale: 'en',
  messages,
})

export default i18n
