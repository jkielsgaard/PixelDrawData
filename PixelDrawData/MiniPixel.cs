using LiteDB;
using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace PixelDrawData
{
    public class MiniPixelJSON_DB
    {
        public int ID { get; set; }
        public string PixelJSON { get; set; }
    }

    public class MiniPixelJSON
    {
        public string PixelID { get; set; }
        public int[,] PixelDataXY { get; set; }
        public string PixelLabel { get; set; }
    }


    public class MiniPixel
    {
        public void run()
        {
            /// TRAINING DATASET
            for (int i = 0; i < 10; i++)
            {
                MiniPixelJSON pd = PixelDataGenerator();
                PixelConsolDrawing(pd);

                string MiniPixelLabel = Console.ReadLine();
                pd.PixelLabel = MiniPixelLabel;
                PutData(JsonConvert.SerializeObject(pd));
            }
            /// 

            /// AI proposals
            for (int i = 0; i < 10; i++)
            {
                MiniPixelJSON pd = PixelDataGenerator();
                string result = AI(pd);
                PixelConsolDrawing(pd);
                Console.WriteLine(result);
                Console.Read();
            }
            ///

            Console.Read();
        }

        private string AI(MiniPixelJSON data)
        {
            double s = 0;
            double c = 0;

            List<MiniPixelJSON_DB> PixelList = GetData();
            foreach (MiniPixelJSON_DB Pixeldb in PixelList)
            {
                MiniPixelJSON Pixel = JsonConvert.DeserializeObject<MiniPixelJSON>(Pixeldb.PixelJSON);

                int xy = 4;

                int compare = 0;

                for (int x = 0; x < xy; x++)
                {
                    for (int y = 0; y < xy; y++)
                    {
                        if (Pixel.PixelDataXY[x, y] == 1)
                        {
                            if (data.PixelDataXY[x, y] == Pixel.PixelDataXY[x, y])
                            {
                                compare++;
                            }
                        }
                    }
                }

                double comparedpercent = (int)Math.Round((double)(100 * compare) / 4);

                //if (comparedpercent == 100)
                //{
                //    string result = "";
                //    if (Pixel.PixelLabel == "s") { result = "S=100% / C=0%"; }
                //    if (Pixel.PixelLabel == "c") { result = "S=0% / C=100%"; }

                //    return result;
                //}

                if (Pixel.PixelLabel == "s") { s = s / comparedpercent; }
                if (Pixel.PixelLabel == "c") { c = c / comparedpercent; }
            }

            //s = s / PixelList.Count;
            //c = c / PixelList.Count;

            return "S=" + s.ToString() + "% / C=" + c.ToString() + "%";
        }

        private void PixelConsolDrawing(MiniPixelJSON data)
        {
            Console.WriteLine("──────────────────────────────");
            Console.WriteLine();
            Console.WriteLine("PixelID:    " + data.PixelID);
            Console.WriteLine("PixelLabel: " + data.PixelLabel);
            Console.WriteLine();


            Console.Write("┌");
            for (int i = 0; i < 4; i++) { Console.Write("─"); }
            Console.WriteLine("┐");

            for (int y = 0; y <= 3; y++)
            {
                Console.Write("│");
                for (int x = 0; x <= 3; x++)
                {
                    if (data.PixelDataXY[y, x] == 0) { Console.Write(" "); }
                    else if (data.PixelDataXY[y, x] == 1) { Console.Write("█"); }
                }
                Console.WriteLine("│");
            }
            Console.Write("└");
            for (int i = 0; i < 4; i++) { Console.Write("─"); }
            Console.WriteLine("┘");
            Console.WriteLine();
        }

        private MiniPixelJSON PixelDataGenerator()
        {
            int[,] PixelDrawXY = new int[4, 4];

            Random r = new Random();

            int z = r.Next(4);
            int d;
            int sd = 0;

            for (int y = 0; y <= 3; y++)
            {
                for (int x = 0; x <= 3; x++)
                {
                    if (x == z) { PixelDrawXY[x, y] = 1; }
                    else { PixelDrawXY[x, y] = 0; }
                }

                d = r.Next(-1, 2);
                if (y != 0)
                {
                    if (d == 1 && sd == -1) { d = r.Next(-1, 1); }
                    else if (d == -1 && sd == 1) { d = r.Next(2); }
                }
                sd = d;
                z = z + d;

                if (z == -1) { z++; }
                if (z == 4) { z--; }
            }

            MiniPixelJSON p = new MiniPixelJSON();

            p.PixelID = HexGen();
            p.PixelLabel = "-";
            p.PixelDataXY = PixelDrawXY;

            return p;
        }

        private string HexGen()
        {
            var r = new Random();
            int A = r.Next(10000000, 999999999);
            return A.ToString("X");
        }

        private void PutData(string data)
        {
            using (var db = new LiteDatabase(@"MiniPixelTrainingData"))
            {
                var pdbs = db.GetCollection<PixelDataJSON>("pdb");

                var pdb = new PixelDataJSON
                {
                    PixelJSON = data
                };

                pdbs.Insert(pdb);
            }
        }

        private List<MiniPixelJSON_DB> GetData()
        {
            using (var db = new LiteDatabase(@"MiniPixelTrainingData"))
            {
                var pdbs = db.GetCollection<MiniPixelJSON_DB>("pdb");

                return pdbs.FindAll().ToList();
            }
        }

    }
}
