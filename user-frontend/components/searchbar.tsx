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

  const retrieveCurrentPincodeDataFromLocalStorage = () => {
    let data: MerchantData[] = JSON.parse(
      window.localStorage.getItem("merchantData") || "[]"
    );
    let currentPincodeData: MerchantData[] = [];
    for (let i = 0; i < data.length; i++) {
      if (data[i].pincodes.includes(pincode)) {
        currentPincodeData.push(data[i]);
      }
    }
    return currentPincodeData;
  };

  const renewData = async () => {
    let data = await fetchPincodeData();
    window.localStorage.setItem("merchantData", JSON.stringify(data));
    let newPincodeData = retrieveCurrentPincodeDataFromLocalStorage();
    setMerchantData(newPincodeData);
  }

  const fetchPincodeData = async () => {
    // TODO: Change the fetch endpoint.
    // const res = await fetch(`http://192.168.180.253:3001/pincode/${pincode}`);
    // const data: MerchantData[] = await res.json();
    // console.log(data)
    // Mock data
    const data = [
      {
        name: "John doe",
        email: "john.doe@gmail.com",
        pincodes: ["121007", "121003", "121004"],
      },
      {
        name: "New",
        email: "newMe@gmail.com",
        pincodes: ["121002"],
      },
      {
        name: "Lol",
        email: "john.doe@gmail.com",
        pincodes: ["121002", "121003", "121004"],
      },
      {
        name: "hehehh",
        email: "newnwenw@gmail.com",
        pincodes: ["121005", "121006"],
      },
      
    ];
    return data;
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    let start = performance.now();
    if (!window.localStorage.getItem("merchantData")) {
      let data = await fetchPincodeData();
      window.localStorage.setItem("merchantData", JSON.stringify(data));
      setMerchantData(data);
    } else {
      let currentPincodeData = retrieveCurrentPincodeDataFromLocalStorage();
      if (currentPincodeData.length === 0) {
        await renewData()
      } else {
        setMerchantData(currentPincodeData);
      }
    }

    let end = performance.now();
    setTimeTakenForRequest(end - start); // in ms

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
              className="bg-sky-50 dark:bg-gray-800 transition-all duration-300 p-6 ps-12 text-md rounded-full focus:shadow-md"
            />
          </form>
        </div>
      </div>
    </>
  );
}
