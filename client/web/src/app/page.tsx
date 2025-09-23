"use client";

import { useChatWebsocket } from "@/shared/api/api.websocket";
import { ChatInput } from "@/shared/chat-input";
import { BaseLayout } from "@/widgets/layouts/base-layout";
import { IconHelicopterLanding } from "@tabler/icons-react";

export default function Home() {
  const { isOpened, sendData } = useChatWebsocket();

  return (
    <BaseLayout>
      {!isOpened ? <IconHelicopterLanding /> : <h1>Connected</h1>}
      <div className="relative">
        <ChatInput onSend={sendData} />
      </div>
    </BaseLayout>
  );
}
