"use client";
import { useEffect, useState } from "react"
import { FNV1a } from "../security/hash";

const baseURL = process.env.NEXT_PUBLIC_WEBSOCKET_URL

export const useChatWebsocket = () => {
    const [isOpened, setIsOpened] = useState(false)
    const [ws, setWs] = useState<WebSocket | null>(null);

    function messageListener(this: WebSocket, ev: MessageEvent<ArrayBuffer>) {
        const uint8 = new Uint8Array(ev.data);
        const view = new DataView(ev.data);

        const data = uint8.subarray(0, uint8.length - 4) // without hashsum

        const computed = FNV1a(data);
        const stored = view.getUint32(uint8.length - 4, true);

        if (computed != stored) {
            console.log("INVALID HASHSUM")
            return;
        }

        console.log("VALID HASHSUM")
    }

    function openListener(this: WebSocket) {
        setIsOpened(true)
    }

    useEffect(() => {
        const socket = new WebSocket(new URL("chat", baseURL))
        socket.binaryType = "arraybuffer";
        setWs(socket)

        socket.addEventListener("message", messageListener)
        socket.addEventListener("open", openListener)

        return () => {
            socket.close()
        }
    }, [])

    const sendData = (data: string | ArrayBufferLike) => {
        if (data instanceof ArrayBuffer) {
            ws?.send(data)
            return;
        }

    }

    return {
        isOpened,
        ws,
        sendData,
    }
}