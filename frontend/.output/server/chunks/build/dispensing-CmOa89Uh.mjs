import { u as useNuxtApp } from './server.mjs';
import { defineComponent, ref, unref, useSSRContext } from 'vue';
import { ssrRenderAttrs, ssrRenderAttr } from 'vue/server-renderer';
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
  __name: "dispensing",
  __ssrInlineRender: true,
  setup(__props) {
    const { $api } = useNuxtApp();
    const visitId = ref("");
    const itemId = ref("");
    const qty = ref(1);
    const ok = ref(false);
    return (_ctx, _push, _parent, _attrs) => {
      _push(`<div${ssrRenderAttrs(_attrs)}><h2>Dispensing</h2><form><input${ssrRenderAttr("value", unref(visitId))} placeholder="Visit ID"><input${ssrRenderAttr("value", unref(itemId))} placeholder="Item ID"><input${ssrRenderAttr("value", unref(qty))} type="number" step="1" placeholder="Qty"><button type="submit">Proses</button></form>`);
      if (unref(ok)) {
        _push(`<div class="ok">OK</div>`);
      } else {
        _push(`<!---->`);
      }
      _push(`</div>`);
    };
  }
});
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/pharmacy/dispensing.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};

export { _sfc_main as default };
//# sourceMappingURL=dispensing-CmOa89Uh.mjs.map
