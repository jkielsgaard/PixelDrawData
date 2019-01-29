using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace PixelDrawData
{
    public class PixelJSON
    {
        public string PixelID { get; set; }
        public int[,] PixelDataXY { get; set; }
        public string PixelLabel { get; set; }
    }

    public class PixelData
    {

        public string PixelDataGenerator()
        {
            int[,] PixelDrawXY = new int[16, 16];

            Random r = new Random();

            int z = r.Next(16);
            int d;
            int sd = 0;

            for (int y = 0; y <= 15; y++)
            {
                for (int x = 0; x <= 15; x++)
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
                if (z == 16) { z--; }
            }

            PixelJSON p = new PixelJSON();

            p.PixelID = HexGen();
            p.PixelLabel = "-";
            p.PixelDataXY = PixelDrawXY;

            return JsonConvert.SerializeObject(p);
        }


        public PixelJSON PixelDataGeneratorJSON()
        {
            int[,] PixelDrawXY = new int[16, 16];

            Random r = new Random();

            int z = r.Next(16);
            int d;
            int sd = 0;

            for (int y = 0; y <= 15; y++)
            {
                for (int x = 0; x <= 15; x++)
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
                if (z == 16) { z--; }
            }


            if (PixelDrawXY[0, 0] == 1 || PixelDrawXY[1, 1] == 1 || PixelDrawXY[1, 0] == 1 || PixelDrawXY[0, 1] == 1)
            {
                PixelDrawXY[0, 0] = 1;
                PixelDrawXY[1, 1] = 1;
                PixelDrawXY[1, 0] = 1;
                PixelDrawXY[0, 1] = 1;
            }

            for (int y = 0; y < 15; y = y + 2)
            {
                for (int x = 0; x < 15; x = x + 2)
                {
                    if (PixelDrawXY[x, y] == 1 || PixelDrawXY[x+1, y+1] == 1 || PixelDrawXY[x+1, y] == 1 || PixelDrawXY[x, y+1] == 1)
                    {
                        PixelDrawXY[x, y] = 1;
                        PixelDrawXY[x+1, y+1] = 1;
                        PixelDrawXY[x+1, y] = 1;
                        PixelDrawXY[x, y+1] = 1;
                    }
                }
            }

            PixelJSON p = new PixelJSON();

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
    }
}
