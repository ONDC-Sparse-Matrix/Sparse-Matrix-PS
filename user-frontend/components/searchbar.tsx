import { Input } from "./ui/input";
import { Search } from "lucide-react";

export function SearchBar() {
  return (
    <>
      <div className="relative w-full mb-10">
        <div className="max-w-xl mx-auto">
          <div className="absolute inset-y-0 flex items-center ps-3 pointer-events-none">
            <Search className="text-sm text-gray-400 dark:text-gray-600" />
          </div>
          <Input placeholder="Search Pincode..." className="bg-gray-50 dark:bg-gray-800 transition-all duration-300 p-6 ps-12 text-md rounded-full focus:shadow-md" />
        </div>
      </div>
    </>
  );
}
