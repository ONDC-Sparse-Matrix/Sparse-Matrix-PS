"use client";

import { Input } from "./ui/input";
import { Search } from "lucide-react";
import { useState } from "react";
import { useRouter } from "next/navigation";

export function SearchBar() {
  const [pincode, setPincode] = useState("");
  const router = useRouter();

  const handlePincodeChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPincode(e.target.value);
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    router.push(`/search?pincode=${pincode}`);
  }

  return (
    <>
      <div className="relative w-full mb-10">
        <div className="max-w-xl mx-auto">
          <div className="absolute inset-y-0 flex items-center ps-3 pointer-events-none">
            <Search className="text-sm text-gray-400 dark:text-gray-600" />
          </div>
          <form onSubmit={handleSubmit}>
            <Input
              placeholder="Enter your pincode"
              onChange={(e) => {
                setPincode(e.target.value);
              }}
              className="bg-gray-50 dark:bg-gray-800 transition-all duration-300 p-6 ps-12 text-md rounded-full focus:shadow-md"
            />
          </form>
        </div>
      </div>
    </>
  );
}
