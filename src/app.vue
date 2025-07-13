<template>
  <div>
    <nav class="navbar navbar-expand-lg bg-body-tertiary">
      <div class="container-fluid">
        <a
          class="navbar-brand"
          href="#"
        >{{ $t('appTitle') }}</a>

        <button
          class="navbar-toggler"
          type="button"
          data-toggle="collapse"
          data-target="#navbarSupportedContent"
          aria-controls="navbarSupportedContent"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span class="navbar-toggler-icon" />
        </button>

        <div
          id="navbarSupportedContent"
          class="collapse navbar-collapse"
        >
          <ul class="navbar-nav me-auto">
            <li class="nav-item">
              <a
                class="nav-link"
                href="https://github.com/Luzifer/password#via-api"
              >{{ $t('menuAPIDocs') }}</a>
            </li>
          </ul>

          <ul class="navbar-nav">
            <li class="nav-item">
              <button
                type="button"
                class="btn btn-default navbar-btn"
                @click="showSettingsModal = true"
              >
                <i
                  class="fas fa-cog"
                  aria-hidden="true"
                /> {{ $t('menuSettings') }}
              </button>
            </li>
          </ul>
        </div>
      </div>
    </nav>

    <div class="container mt-4">
      <div class="row justify-content-md-center">
        <div class="col-xs-12 col-sm-8 col-md-6">
          <div class="p-5 mb-4 bg-body-tertiary rounded-3">
            <p class="fs-5 m-0 text-center fw-light">
              {{ $t('lead', {refreshRate}) }}
            </p>
          </div>

          <div class="card mb-4">
            <div class="card-body">
              <div class="form-group">
                <input
                  id="focusedInput"
                  ref="focusedInput"
                  v-model="password"
                  class="form-control text-center font-monospace border-0"
                  type="text"
                  @focus="focusPassword"
                  @blur="unfocusPassword"
                >
              </div>
            </div>
            <div class="card-footer">
              <div
                class="progress"
                style="height: 5px;"
              >
                <div
                  class="progress-bar"
                  :style="`width: ${progessWidth}%;`"
                />
              </div>
            </div>
          </div>

          <div class="card">
            <div class="card-header">
              {{ $t('faq.title') }}
            </div>
            <div class="card-body">
              <ul>
                <li>
                  <strong>{{ $t('faq.store.title') }}</strong><br>
                  {{ $t('faq.store.content') }}
                </li>
                <li>
                  <strong>{{ $t('faq.secure.title') }}</strong><br>
                  {{ $t('faq.secure.content') }}
                </li>
                <li>
                  <strong>{{ $t('faq.opensource.title') }}</strong><br>
                  {{ $t('faq.opensource.content.pre') }}
                  <a href="https://github.com/Luzifer/password">GitHub</a>
                  {{ $t('faq.opensource.content.post') }}
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>

    <settings-modal
      :show="showSettingsModal"
      @hide="settingsModalHidden"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import SettingsModal from './settings.vue'
import storedSettings from './storage'

const REGENERATE_AFTER = 30000 // milliseconds
const TICK_RATE = 100 // milliseconds

export default defineComponent({
  components: { SettingsModal },

  computed: {
    progessWidth(): number {
      return this.timeLeft / REGENERATE_AFTER * 100
    },

    refreshRate(): String {
      return (REGENERATE_AFTER / 1000).toFixed(0)
    },
  },

  data() {
    return {
      password: '...',
      settings: {} as any,
      showSettingsModal: false,
      ticker: null as number | null,
      timeLeft: REGENERATE_AFTER,
    }
  },

  methods: {
    fetchPassword(): Promise<void> {
      const params = new URLSearchParams({
        length: this.settings.passwordLength,
        separator: this.settings.xkcdSeparator,
        special: this.settings.useSpecial,
        xkcd: this.settings.useXKCD,
      })

      return fetch(`/v1/getPassword?${params.toString()}`)
        .then((resp: Response) => resp.text())
        .then((pass: string) => {
          this.password = pass
          this.timeLeft = REGENERATE_AFTER
        })
    },

    focusPassword(): void {
      if (!this.ticker) {
        return
      }
      window.clearInterval(this.ticker)
      this.ticker = null

      this.$refs.focusedInput.select()
    },

    settingsModalHidden(): void {
      this.showSettingsModal = false
      this.settings = storedSettings.get()
    },

    tick() {
      this.timeLeft -= TICK_RATE
      if (this.timeLeft <= 0) {
        this.fetchPassword()
      }
    },

    unfocusPassword(): void {
      if (this.ticker) {
        return
      }
      // Reset timer to give more space on accidental blur
      this.timeLeft = REGENERATE_AFTER
      this.ticker = window.setInterval(() => this.tick(), TICK_RATE)
    },
  },

  mounted() {
    this.settings = storedSettings.get()
    this.unfocusPassword() // Start ticker
  },

  name: 'PasswordApp',

  watch: {
    settings() {
      this.fetchPassword()
    },
  },
})
</script>
