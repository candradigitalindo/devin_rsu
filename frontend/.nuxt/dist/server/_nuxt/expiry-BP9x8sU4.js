import { u as useNuxtApp } from "../server.mjs";
import { defineComponent, ref, unref, useSSRContext } from "vue";
import { ssrRenderAttrs, ssrRenderAttr, ssrRenderList, ssrInterpolate } from "vue/server-renderer";
import "ofetch";
import "#internal/nuxt/paths";
import "hookable";
import "unctx";
import "h3";
import "unhead";
import "radix3";
const _sfc_main = /* @__PURE__ */ defineComponent({
  __name: "expiry",
  __ssrInlineRender: true,
  setup(__props) {
    const { $api } = useNuxtApp();
    const before = ref("");
    const rows = ref([]);
    return (_ctx, _push, _parent, _attrs) => {
      _push(`<div${ssrRenderAttrs(_attrs)}><h2>Expiry Report</h2><input type="date"${ssrRenderAttr("value", unref(before))}><button>Tampilkan</button><ul><!--[-->`);
      ssrRenderList(unref(rows), (b) => {
        _push(`<li>${ssrInterpolate(b.item_name)} - ${ssrInterpolate(b.batch_no)} - ${ssrInterpolate(b.expiry_date)}</li>`);
      });
      _push(`<!--]--></ul></div>`);
    };
  }
});
const _sfc_setup = _sfc_main.setup;
_sfc_main.setup = (props, ctx) => {
  const ssrContext = useSSRContext();
  (ssrContext.modules || (ssrContext.modules = /* @__PURE__ */ new Set())).add("pages/reports/expiry.vue");
  return _sfc_setup ? _sfc_setup(props, ctx) : void 0;
};
export {
  _sfc_main as default
};
//# sourceMappingURL=expiry-BP9x8sU4.js.map
