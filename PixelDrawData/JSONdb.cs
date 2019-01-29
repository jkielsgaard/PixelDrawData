using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using LiteDB;
using Newtonsoft.Json;

namespace PixelDrawData
{
    public class PixelDataJSON
    {
        public int ID { get; set; }
        public string PixelJSON { get; set; }
    }

    public class JSONdb
    {
        public void PutData(string data)
        {
            using (var db = new LiteDatabase(@"PixelTrainingData"))
            {
                var pdbs = db.GetCollection<PixelDataJSON>("pdb");

                var pdb = new PixelDataJSON { PixelJSON = data };

                pdbs.Insert(pdb);
            }
        }

        public List<PixelDataJSON> GetData()
        {
            using (var db = new LiteDatabase(@"PixelTrainingData"))
            {
                var pdbs = db.GetCollection<PixelDataJSON>("pdb");
                
                return pdbs.FindAll().ToList();
            }
        }
    }
}
