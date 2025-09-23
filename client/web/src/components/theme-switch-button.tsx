"use client";

import { useTheme } from "next-themes";
import { Button } from "./ui/button";
import { IconSun } from "@tabler/icons-react";
import capitalize from "@/shared/utils/capitalize";
import { useEffect, useState } from "react";

export function ThemeSwitchButton() {
  const [mounted, setMounted] = useState<boolean>();
  const { setTheme, theme } = useTheme();

  useEffect(() => {
    setMounted(true);
  }, []);

  if (!mounted) {
    return null;
  }

  const switchThemes = () => {
    setTheme((prev) => {
      switch (prev) {
        case "dark":
          return "light";
        case "light":
          return "system";
        case "system":
          return "dark";
        default:
          return "dark";
      }
    });
  };

  return (
    <Button
      onClick={switchThemes}
      variant="ghost"
      asChild
      size="sm"
      className="hidden sm:flex cursor-pointer"
    >
      <div>
        <IconSun className="!size-4" />
        {capitalize(theme ?? "")}
      </div>
    </Button>
  );
}
