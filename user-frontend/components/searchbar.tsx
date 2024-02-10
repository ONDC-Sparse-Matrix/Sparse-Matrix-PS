"use client";

import { Input } from "./ui/input";
import { Search } from "lucide-react";
import { useState } from "react";
import { useRouter } from "next/navigation";
import { MerchantData } from "@/lib/types";
import { useStore } from "@/lib/store";

interface SearchBarProps {
  pincode?: string;
}

export function SearchBar(props: SearchBarProps) {
  const [pincode, setPincode] = useState(props.pincode ? props.pincode : "");
  const router = useRouter();

  const { setMerchantData, setTimeTakenForRequest } = useStore();

  const fetchPincodeData = async () => {
    // TODO: Change the fetch endpoint.
    // const res = await fetch(`localhost:3000/pincode/${pincode}`);
    // const data: MerchantData[] = await res.json();

    // Mock data
    const data: MerchantData[] = [
      {
        name: "John Doe",
        email: "john.doe@gmail.com",
        pincodes: ["110001", "110002", "110003"],
      },
    ];
    return data;
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    let start = performance.now();
    if (!window.localStorage.getItem("merchantData")) {
      const data = await fetchPincodeData();
      window.localStorage.setItem("merchantData", JSON.stringify(data));
      setMerchantData(data);
    } else {
      let data: MerchantData[] = JSON.parse(
        window.localStorage.getItem("merchantData") || "[]"
      );
      let currentPincodeData: MerchantData[] = [];
      for (let i = 0; i < data.length; i++) {
        if (data[i].pincodes.includes(pincode)) {
          currentPincodeData.push(data[i]);
        }
      }
      setMerchantData(currentPincodeData);
    }

    let end = performance.now();
    setTimeTakenForRequest(end - start);

    router.push(`/search?pincode=${pincode}`);
  };

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
              type="number"
              defaultValue={props.pincode ? props.pincode : ""}
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
