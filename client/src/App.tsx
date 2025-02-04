import { useEffect, useState } from "react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

interface ShortenedURL {
  shortened: string;
  original: string;
  created_at: string;
}

export default function App() {
  const [url, setUrl] = useState("");
  const [loading, setLoading] = useState(false);
  const [shortenedUrls, setShortenedUrls] = useState<ShortenedURL[]>([]);
  const [alert, setAlert] = useState("");

  useEffect(() => {
    fetch("http://localhost:8080/api/urls")
      .then((res) => res.json())
      .then((data) =>
        setShortenedUrls(
          data.urls.map((url: ShortenedURL) => {
            return {
              ...url,
              shortened: `/r/${url.shortened}`,
              created_at: new Date(url.created_at)
                .toLocaleString("pt-BR", {
                  day: "2-digit",
                  month: "2-digit",
                  year: "numeric",
                  hour: "2-digit",
                  minute: "2-digit",
                })
                .replace(",", " às"),
            };
          }),
        ),
      );
  }, []);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    const res = await fetch("http://localhost:8080/api/urls", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ original: url }),
    });

    if (!res.ok) {
      return;
    } // handle error

    const body = await res.json();

    setShortenedUrls([
      {
        shortened: `/r/${body.shortened}`,
        original: body.original,
        created_at: new Date(body.created_at)
          .toLocaleString("pt-BR", {
            day: "2-digit",
            month: "2-digit",
            year: "numeric",
            hour: "2-digit",
            minute: "2-digit",
          })
          .replace(",", " às"),
      },
      ...shortenedUrls,
    ]);

    setAlert("URL: http://localhost:8080/r/" + body.shortened);

    setUrl("");
  };

  return (
    <div className="container mx-auto px-4 py-8 max-w-4xl">
      <div className="flex justify-center mb-12">
        <h1 className="text-4xl font-bold">
          EN<span className="text-[#00BCD4]">CURTA</span>
          <span className="text-[#00BCD4]">GO</span>
        </h1>
      </div>

      <div className="space-y-8">
        <div className="space-y-4">
          <h2 className="text-2xl font-semibold text-center">
            Encurte sua URL
          </h2>
          <form onSubmit={handleSubmit} className="space-y-4">
            <Input
              type="url"
              placeholder="Insira sua URL"
              value={url}
              onChange={(e) => setUrl(e.target.value)}
              required
            />
            <Button
              type="submit"
              className="w-full bg-[#00BCD4] hover:bg-[#00ACC1]"
              disabled={loading}
            >
              Encurta
            </Button>
          </form>
          <div className="text-sm text-center text-gray-600">
            {alert && (
              <a href={alert.split(" ")[1]} target="_blank">
                {alert}
              </a>
            )}
          </div>
        </div>

        <div className="space-y-4">
          <h2 className="text-2xl font-semibold">Últimas URLs encurtadas</h2>
          <div className="border rounded-lg">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>URL</TableHead>
                  <TableHead>ORIGINAL</TableHead>
                  <TableHead>Criado em</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {shortenedUrls.length > 0 ? (
                  shortenedUrls.map((item, index) => (
                    <TableRow key={index}>
                      <TableCell className="font-medium">
                        {item.shortened}
                      </TableCell>
                      <TableCell>{item.original}</TableCell>
                      <TableCell>{item.created_at}</TableCell>
                    </TableRow>
                  ))
                ) : (
                  <TableRow>
                    <TableCell colSpan={3} className="text-center">
                      Nenhuma URL encurtada
                    </TableCell>
                  </TableRow>
                )}
              </TableBody>
            </Table>
          </div>
        </div>
      </div>
    </div>
  );
}
