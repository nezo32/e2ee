"use client";

import { IconSend2 } from "@tabler/icons-react";
import {
  ChangeEventHandler,
  KeyboardEventHandler,
  useRef,
  useState,
} from "react";

export type ChatInputProps = {
  onSend?: (text: string, files?: File[]) => Promise<void> | void;
};

export function ChatInput({ onSend }: ChatInputProps) {
  const textareaRef = useRef<HTMLTextAreaElement | null>(null);
  const [text, setText] = useState("");

  const keydownHandler: KeyboardEventHandler<HTMLTextAreaElement> = (e) => {
    const enter = e.key === "Enter";
    const shift = e.shiftKey;

    if (enter && !shift) {
      e.preventDefault();

      onSend?.(text);
      return;
    }
  };

  const changeHandler: ChangeEventHandler<HTMLTextAreaElement> = (e) => {
    setText(e.target.value);
  };

  return (
    <div className="absolute w-full h-[calc(100vh-72px)]">
      <div className="absolute w-full bottom-0 p-4 py-3 flex gap-2 bg-secondary border-t">
        <textarea
          ref={textareaRef}
          onKeyDown={keydownHandler}
          onChange={changeHandler}
          placeholder="Write a message..."
          className="resize-none w-full max-h-50 scroll-none self-center outline-none text-base/5 field-sizing-content"
        />
        <div className="shrink-0 self-end">
          <IconSend2 className="!size-6 cursor-pointer text-secondary-foreground" />
        </div>
      </div>
    </div>
  );
}
