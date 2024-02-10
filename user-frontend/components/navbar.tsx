import { Button } from "./ui/button";
import { Github } from "lucide-react";
import { Separator } from "@/components/ui/separator";

export function Navbar() {
  return (
    <>
      <nav className="bg-transparent w-full fixed backdrop-blur-lg dark:bg-gray-900">
        <div className="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
          <a
            href="/"
            className="flex items-center space-x-3 rtl:space-x-reverse"
          >
            <span className="text-2xl font-extrabold text-blue-500 whitespace-nowrap dark:text-white">
              matriX
            </span>
          </a>
          <div className="hidden w-full md:block md:w-auto" id="navbar-default">
            <ul className="p-4 md:p-0 mt-4 rounded-lg bg-gray-50 md:flex-row md:space-x-8 rtl:space-x-reverse md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700">
              <li>
                <a href="https://github.com/ONDC-Sparse-Matrix" target="_blank">
                  <Button variant={"link"}>
                    <Github />
                  </Button>
                </a>
              </li>
            </ul>
          </div>
        </div>
      </nav>
    </>
  );
}
