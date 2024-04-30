import { LaptopIcon } from "@radix-ui/react-icons";
import {
  AppWindowIcon,
  InfoIcon,
  Settings2Icon,
  SquareTerminal
} from "lucide-react";
import Image from "next/image";
import Link from "next/link";

export default function Sidebar() {
  const icon_size = {
    h: 24,
    w: 24
  };

  return (
    <div className="min-h-screen py-4 bg-black flex flex-col items-center border-r">
      <ul className="flex flex-col items-center space-y-4 sticky top-4">
        <li className="relative w-10 h-10 hover:cursor-pointer gen-transition">
          <Link href="/">
            <Image
              src="/logo.png"
              fill={true}
              alt="logo"
              objectPosition="center"
              className="rounded-full left-0 right-0 ml-auto mr-auto border hover:border-foreground gen-transition p-1"
            ></Image>
          </Link>
        </li>
        <li className="gen-transition border rounded-full hover:border-foreground p-2 text-muted hover:text-foreground hover:cursor-pointer">
          <Link href="/console">
            <AppWindowIcon
              width={icon_size.w}
              height={icon_size.h}
            ></AppWindowIcon>
          </Link>
        </li>
        <li className="gen-transition border rounded-full hover:border-foreground p-2 text-muted hover:text-foreground hover:cursor-pointer">
          <Link href="/settings">
            <Settings2Icon
              width={icon_size.w}
              height={icon_size.h}
            ></Settings2Icon>
          </Link>
        </li>
      </ul>
    </div>
  );
}
