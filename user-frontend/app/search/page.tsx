'use client';

import { useSearchParams } from "next/navigation";
import { MerchantCard } from "@/components/merchant-card";

export default function Search() {
  const searchParams = useSearchParams();
  const pincode = searchParams.get("pincode");

  return (
    <div className="h-screen max-w-xl mx-auto pt-20">
      <h1 className="text-lg">Search Results for {pincode}</h1>
      <MerchantCard />
    </div>
  );
}
