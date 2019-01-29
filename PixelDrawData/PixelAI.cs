using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace PixelDrawData
{
    public class PixelAI
    {
        public PixelJSON AI(PixelJSON PixelData)
        {
            int xyLength = 15;

            string[] xy = new string[xyLength];

            int c = 0;
            for (int x = 0; x < xyLength; x = x + 2)
            {
                for (int y = 0; y < xyLength; y = y + 2)
                {
                    if (PixelData.PixelDataXY[y, x] == 1)
                    {
                        int x1 = x;
                        int y1 = y;

                        if (x1 == 2) { x1 = 1; }
                        if (x1 == 4) { x1 = 2; }
                        if (x1 == 6) { x1 = 3; }
                        if (x1 == 8) { x1 = 4; }
                        if (x1 == 10) { x1 = 5; }
                        if (x1 == 12) { x1 = 6; }
                        if (x1 == 14) { x1 = 7; }
                        if (y1 == 2) { y1 = 1; }
                        if (y1 == 4) { y1 = 2; }
                        if (y1 == 6) { y1 = 3; }
                        if (y1 == 8) { y1 = 4; }
                        if (y1 == 10) { y1 = 5; }
                        if (y1 == 12) { y1 = 6; }
                        if (y1 == 14) { y1 = 7; }

                        xy[c] = y1 + "," + x1;
                        c++;
                    }
                }
            }

            for (int i = 0; i < xyLength; i++) { Console.WriteLine(xy[i]); } //only to see the coordinates

            int downCount = 0;
            int upCount = 0;

            for (int i = 1; i < xyLength; i++)
            {
                if (xy[i] == null) { break; }

                string[] xyFirst = xy[i-1].Split(',');
                string[] xySecond = xy[i].Split(',');

                if (int.Parse(xyFirst[0]) < int.Parse(xySecond[0])) { downCount++; }
                if (int.Parse(xyFirst[0]) > int.Parse(xySecond[0])) { upCount++; }
            }

            if ((upCount == 0 || downCount == 0) || (downCount == 0 && upCount == 0))
            {
                Console.WriteLine("S");
            }

            return PixelData;
        }

        public PixelJSON ML(PixelJSON PixelData) { return PixelData; }
    }
}
