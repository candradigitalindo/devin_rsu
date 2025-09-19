import { u as useNuxtApp } from './server.mjs';
import { defineComponent, ref, unref, useSSRContext } from 'vue';
import { ssrRenderAttrs, ssrRenderAttr, ssrIncludeBooleanAttr, ssrLooseContain, ssrLooseEqual, ssrRenderList, ssrInterpolate, ssrRenderClass } from 'vue/server-renderer';
import '../nitro/nitro.mjs';
import 'node:http';
import 'node:https';
import 'node:events';
import 'node:buffer';
import 'node:fs';
import 'node:path';
import 'node:crypto';
import 'node:url';
import '../routes/renderer.mjs';
import 'vue-bundle-renderer/runtime';
import 'devalue';
import '@unhead/ssr';
import 'unhead';
import '@unhead/shared';

const _sfc_main = /* @__PURE__ */ defineComponent({
  __name: "index",
  __ssrInlineRender: true,
  setup(__props) {
    const { $api } = useNuxtApp();
    const items = ref([]);
    const expired = ref("");
    const q = ref("");
    function badgeClass(it) {
      return "badge";
    }
    return (_ctx, _push, _parent, _attrs) => {
      _push(`<div${ssrRenderAttrs(_attrs)}><h2>Farmasi - Stok</h2><div class="filters"><select><option${ssrRenderAttr("value", "")}${ssrIncludeBooleanAttr(Array.isArray(unref(expired)) ? ssrLooseContain(unref(expired), "") : ssrLooseEqual(unref(expired), "")) ? " selected" : ""}>Semua</option><option value="true"${ssrIncludeBooleanAttr(Array.isArray(unref(expired)) ? ssrLooseContain(unref(expired), "true") : ssrLooseEqual(unref(expired), "true")) ? " selected" : ""}>Expired</option><option value="false"${ssrIncludeBooleanAttr(Array.isArray(unref(expired)) ? ssrLooseContain(unref(expired), "false") : ssrLooseEqual(unref(expired), "false")) ? " selected" : ""}>Non-Expired</option></select><input${ssrRenderAttr("value", unref(q))} placeholder="Cari item"><button>Terapkan</button></div><table><thead><tr><th>Nama</th><th>Kategori</th><th>Stok</th><th>Status</th></tr></thead><tbody><!--[-->`);
      ssrRenderList(unref(items), (it) => {
        _push(`<tr><td>${ssrInterpolate(it.name)}</td><td>${ssrInterpolate(it.category_name || "-")}</td><td>${ssrInterpolate(it.stock || 0)}</td><td><span class="${ssrRenderClass(badgeClass())}">${ssrInterpolate(it.expiryStatus || "-")}</span></td></tr>`);
      });
      _push(`<!--]--></tbody></table></div>`);
    };
  }
});
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/pharmacy/index.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};

export { _sfc_main as default };
//# sourceMappingURL=index-BklAYvg9.mjs.map
