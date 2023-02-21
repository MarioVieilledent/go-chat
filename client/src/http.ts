import { fetching } from "./storage";

export function get(url: string, script: (res) => void) {
    fetching.set(true);
    fetch(url).then(
        (res) => {
            fetching.set(false);
            script(res);
        }
    )
}
export function getJSON(url: string, script: (res) => void) {
    fetching.set(true);
    fetch(url).then(res => res.json()).then(
        res => {
            fetching.set(false);
            script(res);
        }
    )
}

export function put(url: string, body: string, script: (res) => void) {
    fetching.set(true);
    fetch(url, {
        method: "PUT",
        body: body,
    }).then((res) => {
        fetching.set(false);
        script(res);
    });
}

export function putJSON(url: string, body: string, script: (res) => void) {
    fetching.set(true);
    fetch(url, {
        method: "PUT",
        body: body,
    }).then(res => res.json()).then((res) => {
        fetching.set(false);
        script(res);
    });
}