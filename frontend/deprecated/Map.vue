<template>
    <div>
        <l-map style="height: 800px" :zoom="zoom" :center="center">
            <l-tile-layer :url="url" :attribution="attribution"></l-tile-layer>
            <l-marker :lat-lng="markerLatLng"></l-marker>
        </l-map>
    </div>
</template>

<script>
import L from "leaflet";
import { LMap, LTileLayer, LMarker } from "vue2-leaflet";

export default {
    components: { LMap, LTileLayer, LMarker },
    name: "Map",
    data() {
        return {
            zoom: 16,
            center: [31.020697, 121.426841],
            url: "http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png",
            attribution: "Map data &copy; OpenStreetMap contributors",
            layer: new L.TileLayer("http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"),
            options: {
                position: "bottomleft",
                width: 200,
                height: 175,
            },
        };
    },
    mounted: function () {
        this.setGlobalTitle();
        this.initMap();
    },
    methods: {
        setGlobalTitle: function () {
            var myElement = document.getElementById("global-title");
            myElement.textContent = "Map";
        },
        initMap() {
            let map = L.map("map", {
                minZoom: 3,
                maxZoom: 14,
                center: [39.550339, 116.114129],
                zoom: 12,
                zoomControl: false,
                attributionControl: false,
                crs: L.CRS.EPSG3857,
            });
            this.map = map; //data上需要挂载
            L.tiledLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png").addTo(map);
            L.control
                .zoom({
                    zoomInTitle: "放大",
                    zoomOutTitle: "缩小",
                })
                .addTo(map);
        },
    },
};
</script>

<style scoped>
#map {
    margin: 0;
    overflow: hidden;
    background: #ffffff;
    width: 100%;
    height: 100%;
    position: absolute;
    top: 0;
}
</style>
