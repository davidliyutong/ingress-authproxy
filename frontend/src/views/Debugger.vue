<template>
    <div>
        <v-card class="mx-auto" max-width="800px">
            <div id="terminal" ref="terminal" max-width="800px"></div>

            <v-card-title>Interact with IoT devices via websocket </v-card-title>

            <v-card-subtitle> Configure IP address and port below </v-card-subtitle>
            <v-form>
                <v-container>
                    <v-col cols="12" sm="6">
                        <v-text-field v-model="address" label="Address" :rules="[rules.required, rules.address]" clearable></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6">
                        <v-text-field v-model="port" label="Port" :rules="[rules.required]" clearable maxlength="5"></v-text-field>
                    </v-col>
                </v-container>
            </v-form>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn large color="red lighten-2" text> Clear </v-btn>
                <v-col cols="auto">
                    <v-btn class="mx-2" fab dark right color="green">
                        <v-icon dark> mdi-link </v-icon>
                    </v-btn>
                </v-col>
            </v-card-actions>
            <v-spacer></v-spacer>
        </v-card>
        <v-card> </v-card>
    </div>
</template>
<script>
import { Terminal } from "xterm";
import { FitAddon } from "xterm-addon-fit";
import "xterm/css/xterm.css";
export default {
    data() {
        return {
            term: "", // 保存terminal实例
            rows: 20,
            cols: 80,
            show: false,
            rules: {
                required: (value) => !!value || "Required.",
                counter: (value) => value.length <= 20 || "Max 20 characters",
                email: (value) => {
                    const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
                    return pattern.test(value) || "Invalid e-mail.";
                },
                address: (address) => {
                    const pattern = /^((25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\.){3}(25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))$/;
                    return pattern.test(address) || "Invalid address";
                },
            },
        };
    },
    mounted() {
        this.setGlobalTitle();
        this.initXterm();
    },
    methods: {
        setGlobalTitle: function () {
            var myElement = document.getElementById("global-title");
            myElement.textContent = "Debugger";
        },
        initXterm() {
            let _this = this;
            let term = new Terminal({
                rendererType: "canvas", //渲染类型
                rows: _this.rows, //行数
                cols: _this.cols, // 不指定行数，自动回车后光标从下一行开始
                convertEol: true, //启用时，光标将设置为下一行的开头
                // scrollback: 50, //终端中的回滚量
                disableStdin: false, //是否应禁用输入
                // cursorStyle: "underline", //光标样式
                cursorBlink: true, //光标闪烁
                theme: {
                    foreground: "#ECECEC", //字体
                    background: "#000000", //背景色
                    cursor: "#FFFFFF", //设置光标
                    lineHeight: 20,
                },
            });
            // 创建terminal实例
            term.open(this.$refs["terminal"]);
            // 换行并输入起始符 $
            term.prompt = (_) => {
                term.write("\t\r\n\x1b[36m➜\x1b[0m ");
            };
            term.prompt();
            // canvas背景全屏
            const fitAddon = new FitAddon();
            term.loadAddon(fitAddon);
            fitAddon.fit();

            window.addEventListener("resize", resizeScreen);
            function resizeScreen() {
                try {
                    // 窗口大小改变时，触发xterm的resize方法使自适应
                    fitAddon.fit();
                } catch (e) {
                    console.log("e", e.message);
                }
            }
            _this.term = term;
            _this.runFakeTerminal();
        },
        runFakeTerminal() {
            let _this = this;
            let term = _this.term;
            if (term._initialized) return;
            // 初始化
            term._initialized = true;
            term.writeln("Welcome to \x1b[1;34m Debugger\x1b[0m.");
            term.writeln("This is Web Terminal");
            term.prompt();
            // 添加事件监听器，支持输入方法
            term.onKey((e) => {
                const printable = !e.domEvent.altKey && !e.domEvent.altGraphKey && !e.domEvent.ctrlKey && !e.domEvent.metaKey;
                if (e.domEvent.keyCode === 13) {
                    term.prompt();
                } else if (e.domEvent.keyCode === 8) {
                    // back 删除的情况
                    if (term._core.buffer.x > 2) {
                        term.write("\b \b");
                    }
                } else if (printable) {
                    term.write(e.key);
                }
                console.log(1, "print", e.key);
            });
            term.onData((key) => {
                // 粘贴的情况
                if (key.length > 1) term.write(key);
            });
        },
    },
};
</script>
<style>
#lateral .v-btn--connect {
    bottom: 0;
    position: absolute;
    margin: 32px 32px 32px 32px;
}
</style>
