<script>
import Vue from "vue";

export default Vue.component("image-icon", {
    template: `<v-img :src="imgsrc" max-width="32"></v-img>`,
    props: ["imgsrc"],
});
</script>
