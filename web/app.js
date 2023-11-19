import Vue from 'vue'

export default Vue.createApp({
    setup(props, context) {
        const sub = Vue.ref('');
        const newsub = Vue.ref('');
        const include = Vue.ref('');
        const exclude = Vue.ref('');
        const config = Vue.ref('loading...');
        const checkbox = Vue.ref(false)
        const configurl = Vue.ref('');
        const saveType = Vue.ref(true);
        const saveTypeText = Vue.ref('Parameters are saved within the link');
        const customValue = Vue.ref({
            type: "urltest"
        });
        const customTables = Vue.ref([]);
        const inFetch = Vue.ref(false)
        const inputRef = Vue.ref(null)
        const addTag = Vue.ref(false)

        Vue.watch(checkbox, v => {
            if (v) {
                config.value = oldConfig
            } else {
                configurl.value = ""
            }
        })

        function encryptData(data) {
            var key = CryptoJS.enc.Utf8.parse("your-secret-key");  // 替换为您的密钥
            var encrypted = CryptoJS.AES.encrypt(data, key, {
                mode: CryptoJS.mode.ECB,
                padding: CryptoJS.pad.Pkcs7
            });
            return encrypted.toString();
        }

        let oldConfig = "";

        function saveParameter() {
            const subUrl = new URL(new URL(location.href).origin)
            subUrl.pathname = "/sub"
            const c = config.value != oldConfig ? config.value : ""
            if (c != "") {
                const compressed = pako.deflate(config.value);
                const base64String = Base64.fromUint8Array(compressed, true)
                subUrl.searchParams.set("config", base64String)
            }
            // const ct = customTables.value.length != 0 ? JSON.stringify(customTables.value) : ""
            // if (ct != "") {
            //     const compressed = pako.deflate(ct);
            //     const base64String = Base64.fromUint8Array(compressed, true)
            //     subUrl.searchParams.set("urltest", base64String)
            // }
            configurl.value && subUrl.searchParams.set("configurl", encryptData(configurl.value))
            // include.value && subUrl.searchParams.set("include", include.value)
            // exclude.value && subUrl.searchParams.set("exclude", exclude.value)
            addTag.value && subUrl.searchParams.set("addTag", "true")
            subUrl.searchParams.set("sub", sub.value)
            return subUrl.toString()
        }

        function catchSome(f, onfail) {
            const nf = async (...a) => {
                try {
                    return await f(...a);
                } catch (e) {
                    if (onfail) {
                        onfail()
                    }
                    console.warn(e)
                    alert(String(e))
                }
            }
            return nf
        }



        const click = catchSome(async () => {
            if (sub.value == "") {
                return ""
            }
            if (inFetch.value) {
                return
            }
            newsub.value = ""
            inFetch.value = true
            const subURL = await (async () => {
                // if (saveType.value) {
                return saveParameter()
                // } else {
                //     return await saveServer()
                // }
            })()
            const f = await fetch(subURL)
            if (!f.ok) {
                const msg = await f.text()
                newsub.value = msg
                console.warn(msg)
                inFetch.value = false
                alert("Error: " + msg)
                return
            }
            inFetch.value = false
            newsub.value = subURL
            await Vue.nextTick()
            inputRef.value.scrollIntoView({ behavior: "smooth" })
            inputRef.value.select()
            document.execCommand('copy', true);
            alert("Copied to paste board")
        }, () => {
            inFetch.value = false
        })

        function addUrlTables(v) {
            customTables.value.push(Object.assign({}, Vue.toRaw(v)))
        }

        return {
            sub,
            config,
            include,
            exclude,
            newsub,
            click,
            checkbox,
            configurl,
            saveType,
            saveTypeText,
            customValue,
            customTables,
            inFetch,
            inputRef,
            addTag
        }

    },
}).mount('#app');