import type { ComputedRef, MaybeRef } from 'vue'
export type LayoutKey = string
declare module "../../node_modules/.pnpm/nuxt@3.12.0_@parcel+watcher@2.5.1_db0@0.3.2_eslint@8.57.1_ioredis@5.7.0_magicast@0.3.5_option_th57qugdqeqofwiitkkkl5lhiu/node_modules/nuxt/dist/pages/runtime/composables" {
  interface PageMeta {
    layout?: MaybeRef<LayoutKey | false> | ComputedRef<LayoutKey | false>
  }
}