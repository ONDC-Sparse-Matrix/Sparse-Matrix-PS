import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

export function MerchantCard() {
  return (
    <>
      <Card className="drop-shadow-sm my-4 px-2 transition-all duration-200 rounded-full hover:shadow-md">
        <CardHeader className="flex flex-row justify-between">
          <CardTitle>John Doe</CardTitle>
          <CardDescription>john.doe@gmail.com</CardDescription>
        </CardHeader>
      </Card>
    </>
  );
}
