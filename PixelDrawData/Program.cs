using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace PixelDrawData
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WindowHeight = 40;
            Console.WindowWidth = 40;
            Console.BufferHeight = 1000;
            Console.BufferWidth = 40;

            Program p = new Program();
            PixelData pd = new PixelData();
            JSONdb jdb = new JSONdb();
            PixelAI ai = new PixelAI();

            /// TESTING SMALLSCALE VERSION
            //MiniPixel mini = new MiniPixel();
            //mini.run();
            ///////////////////////////////


            /// TESTING AI FUNCTION
            //PixelJSON data = pd.PixelDataGeneratorJSON();

            //p.PixelConsolAnalyze(JsonConvert.SerializeObject(data));
            //ai.AI(data);

            //for (int i = 0; i < 250; i++)
            //{
            //    p.PixelConsolAnalyze(JsonConvert.SerializeObject(pd.PixelDataGeneratorJSON()));
            //    ai.AI(data);
            //    //Thread.Sleep(1500);
            //}
            ///////////////////////////////

            int answer = 0;
            while (answer != 9)
            {
                PixelJSON pJSON = pd.PixelDataGeneratorJSON();

                p.PixelConsolAnalyze(JsonConvert.SerializeObject(pJSON));

                Console.WriteLine();
                Console.WriteLine("What is the line?");
                Console.WriteLine("1. straight");
                Console.WriteLine("2. curved");
                Console.WriteLine("3. sinewave");
                Console.Write("# ");
                answer = Int32.Parse(Console.ReadLine());

                if (answer == 1) { pJSON.PixelLabel = "straight"; }
                if (answer == 2) { pJSON.PixelLabel = "curved"; }
                if (answer == 3) { pJSON.PixelLabel = "sinewave"; }
                if (answer == 1 || answer == 2 || answer == 3) { jdb.PutData(JsonConvert.SerializeObject(pJSON)); }

            }

            List<PixelDataJSON> PixelDataList = new List<PixelDataJSON>();

            PixelDataList = jdb.GetData();

            foreach (PixelDataJSON item in PixelDataList)
            {
                PixelJSON pj = JsonConvert.DeserializeObject<PixelJSON>(item.PixelJSON);
                Console.WriteLine(pj.PixelID);
            }

            foreach (PixelDataJSON item in PixelDataList)
            {
                p.PixelConsolAnalyze(item.PixelJSON);
                Console.Read();
            }

            Console.WriteLine("Done");
            Console.Read();
        }


        private void PixelConsolDrawing(string JSONdata)
        {
            PixelJSON data = JsonConvert.DeserializeObject<PixelJSON>(JSONdata);

            Console.WriteLine("──────────────────────────────");
            Console.WriteLine();
            Console.WriteLine("PixelID:    " + data.PixelID);
            Console.WriteLine("PixelLabel: " + data.PixelLabel);
            Console.WriteLine();


            Console.Write("┌");
            for (int i = 0; i < 16; i++) { Console.Write("─"); }
            Console.WriteLine("┐");

            for (int y = 0; y <= 15; y++)
            {
                Console.Write("│");
                for (int x = 0; x <= 15; x++)
                {
                    if (data.PixelDataXY[y, x] == 0) { Console.Write(" "); }
                    else if (data.PixelDataXY[y, x] == 1) { Console.Write("█"); }
                }
                Console.WriteLine("│");
            }
            Console.Write("└");
            for (int i = 0; i < 16; i++) { Console.Write("─"); }
            Console.WriteLine("┘");
            Console.WriteLine();
        }

        private void PixelConsolAnalyze(string JSONdata)
        {
            PixelJSON data = JsonConvert.DeserializeObject<PixelJSON>(JSONdata);

            Console.WriteLine("──────────────────────────────");
            Console.WriteLine();
            Console.WriteLine("PixelID:    " + data.PixelID);
            Console.WriteLine("PixelLabel: " + data.PixelLabel);
            Console.WriteLine();


            Console.Write("┌");
            for (int i = 0; i < 23; i++)
            {
                if (i == 2 || i == 5 || i == 8 || i == 11 || i == 14 || i == 17 || i == 20) { Console.Write("┬"); }
                else if (i == 9) { Console.Write("Y"); }
                else if (i == 10) { Console.Write("▼"); }
                else if (i == 12) { Console.Write("X"); }
                else if (i == 13) { Console.Write("►"); }
                else { Console.Write("─"); }
            }
            Console.WriteLine("┐");

            for (int y = 0; y <= 15; y++)
            {
                if (y == 2 || y == 4 || y == 6 || y == 8 || y == 10 || y == 12 || y == 14)
                {
                    Console.WriteLine("├──┼──┼──┼──┼──┼──┼──┼──┤");
                }

                Console.Write("│");
                for (int x = 0; x <= 15; x++)
                {
                    if (data.PixelDataXY[y, x] == 0) { Console.Write(" "); }
                    else if (data.PixelDataXY[y, x] == 1) { Console.Write("█"); }
                    if (x == 1 || x == 3 || x == 5 || x == 7 || x == 9 || x == 11 || x == 13) { Console.Write("│"); }
                }
                Console.WriteLine("│");
            }
            Console.Write("└");
            for (int i = 0; i < 23; i++)
            {
                if (i == 2 || i == 5 || i == 8 || i == 11 || i == 14 || i == 17 || i == 20) { Console.Write("┴"); }
                else { Console.Write("─"); }
            }
            Console.WriteLine("┘");
            Console.WriteLine();
        }
    }



}