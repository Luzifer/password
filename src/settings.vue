<template>
  <div
    id="settingsModal"
    ref="settingsModal"
    class="modal"
    role="dialog"
  >
    <div
      class="modal-dialog"
      role="document"
    >
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">
            {{ $t('menuSettings') }}
          </h5>
          <button
            type="button"
            class="btn-close"
            @click="$emit('hide')"
          />
        </div>
        <div class="modal-body">
          <div class="mb-3">
            <label
              for="passwordLengthOption"
              class="control-label"
            >{{ $t('settings.labelPasswordLength') }}</label>
            <input
              id="passwordLengthOption"
              v-model.number="settings.passwordLength"
              type="number"
              class="form-control"
              max="128"
              min="4"
              step="1"
            >
          </div>
          <div class="mb-3">
            <div class="form-check form-switch">
              <input
                id="useSpecialOption"
                v-model="settings.useSpecial"
                type="checkbox"
                class="form-check-input"
              >
              <label class="form-check-label">{{ $t('settings.labelUseSpecial') }}</label>
            </div>
          </div>
          <div class="mb-3">
            <div class="form-check form-switch">
              <input
                id="useXKCDOption"
                v-model="settings.useXKCD"
                type="checkbox"
                class="form-check-input"
              >
              <label class="form-check-label">
                {{ $t('settings.labelXKCD.pre') }} <a
                  href="https://xkcd.com/936/"
                  target="_blank"
                  rel="noopener noreferrer"
                >{{ $t('settings.labelXKCD.link') }}</a>
              </label>
            </div>
          </div>
          <div class="mb-3">
            <label
              for="xkcdSeparator"
              class="control-label"
            >{{ $t('settings.labelXKCDSeparator') }}</label>
            <input
              id="xkcdSeparator"
              v-model="settings.xkcdSeparator"
              type="text"
              class="form-control"
              placeholder="Optional, could be '-' or any other separator"
            >
          </div>
        </div>
        <div class="modal-footer">
          <button
            type="button"
            class="btn btn-default"
            @click="$emit('hide')"
          >
            {{ $t('settings.btnClose') }}
          </button>
          <button
            id="optionSave"
            type="button"
            class="btn btn-primary"
            @click="saveSettings"
          >
            {{ $t('settings.btnSave') }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { Modal } from 'bootstrap'
import storedSettings from './storage'

export default defineComponent({
  data() {
    return {
      modal: null as Modal | null,
      settings: {} as any,
    }
  },

  emits: ['hide'],

  methods: {
    saveSettings(): void {
      storedSettings.set(this.settings)
      this.$emit('hide')
    },
  },

  mounted() {
    this.modal = new Modal(this.$refs.settingsModal)
    this.$refs.settingsModal.addEventListener('hide.bs.modal', () => this.$emit('hide'))
  },

  name: 'PasswordAppSettings',

  props: {
    show: {
      default: false,
      type: Boolean,
    },
  },

  watch: {
    show(to) {
      if (to) {
        // Reload settings on show in case they changed in another instance
        this.settings = storedSettings.get()
        this.modal.show()
      } else {
        this.modal.hide()
      }
    },
  },
})
</script>
