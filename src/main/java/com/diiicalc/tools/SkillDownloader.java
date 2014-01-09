package com.diiicalc.tools;

import com.diiicalc.api.BlizzardApiError;
import com.diiicalc.core.ActiveSkills;
import com.diiicalc.core.BattlenetRealm;
import com.diiicalc.core.Constants;
import com.diiicalc.core.PassiveSkills;
import org.apache.http.HttpEntity;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.client.utils.URIBuilder;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.util.EntityUtils;
import org.w3c.dom.Document;

import javax.xml.parsers.DocumentBuilder;
import javax.xml.parsers.DocumentBuilderFactory;
import java.io.ByteArrayInputStream;
import java.io.FileOutputStream;
import java.io.OutputStreamWriter;
import java.net.URI;
import java.util.Map;

public class SkillDownloader
{
   public static void main(String[] args) throws Exception
   {
      //downloadPassive();
      downloadActive();
   }

   private static void downloadPassive() throws Exception
   {
      FileOutputStream fos = new FileOutputStream("/Users/barmistead/Desktop/generate.txt");
      OutputStreamWriter writer = new OutputStreamWriter(fos);

      CloseableHttpClient httpClient = HttpClients.createDefault();
      DocumentBuilderFactory dbFactory = DocumentBuilderFactory.newInstance();
      DocumentBuilder dBuilder = dbFactory.newDocumentBuilder();


      for (Map.Entry<String, String[]> entry : PassiveSkills.HERO_TYPE_TO_SLUGS_MAP.entrySet())
      {
         for (String activeSlug : entry.getValue())
         {
            URIBuilder builder = new URIBuilder("http://us.battle.net");

            builder.setPath("/d3/en/tooltip/skill/"+entry.getKey()+"/" + activeSlug);

               URI uri = builder.build();

               HttpGet httpget = new HttpGet(uri);

               CloseableHttpResponse response = httpClient.execute(httpget);

               try
               {
                  HttpEntity entity = response.getEntity();

                  String entityString = EntityUtils.toString(entity);

                  entityString = entityString.replaceAll("\r\n", "");
                  entityString = entityString.replaceAll("\t", "");

                  Document doc = dBuilder.parse(new ByteArrayInputStream(entityString.getBytes("UTF-8")));

                  doc.getDocumentElement().normalize();

                  String skillName = doc.getDocumentElement().getFirstChild().getFirstChild().getTextContent();

                  String style = doc.getDocumentElement().getChildNodes().item(1).getFirstChild().getAttributes().getNamedItem("style").getTextContent();

                  String pattern = ".*background-image: url\\('.*/(.*).png'\\);.*";

                  String icon = style.replaceAll(pattern, "$1");

                  writer.write("SKILL_MAP.put(\"" + activeSlug + "\", new Skill(\"" + activeSlug + "\", \"" + icon + "\", \"" + skillName + "\"));\n");

               }
               catch (Exception ex)
               {
                  ex.printStackTrace();
                  throw ex;
               }
               finally
               {
                  response.close();
               }
            }

      }

      writer.close();
   }

   private static void downloadActive() throws Exception
   {
      String[] runetypes = {null, "a", "b", "c", "d", "e"};

      FileOutputStream fos = new FileOutputStream("/Users/barmistead/Desktop/generate.txt");
      OutputStreamWriter writer = new OutputStreamWriter(fos);

      CloseableHttpClient httpClient = HttpClients.createDefault();
      DocumentBuilderFactory dbFactory = DocumentBuilderFactory.newInstance();
      DocumentBuilder dBuilder = dbFactory.newDocumentBuilder();

      for (Map.Entry<String, String[]> entry : ActiveSkills.HERO_TYPE_TO_SLUGS_MAP.entrySet())
      {
         for (String activeSlug : entry.getValue())
         {
            URIBuilder builder = new URIBuilder("http://us.battle.net");

            builder.setPath("/d3/en/tooltip/skill/"+entry.getKey()+"/" + activeSlug);

            for (String runeType : runetypes)
            {
               if (runeType != null)
               {
                  builder.setParameter("runeType", runeType);
               }

               URI uri = builder.build();

               HttpGet httpget = new HttpGet(uri);

               CloseableHttpResponse response = httpClient.execute(httpget);

               try
               {
                  HttpEntity entity = response.getEntity();

                  String entityString = EntityUtils.toString(entity);

                  entityString = entityString.replaceAll("\r\n", "");
                  entityString = entityString.replaceAll("\t", "");

                  Document doc = dBuilder.parse(new ByteArrayInputStream(entityString.getBytes("UTF-8")));

                  doc.getDocumentElement().normalize();

                  String skillName = doc.getDocumentElement().getFirstChild().getFirstChild().getTextContent();

                  String style = doc.getDocumentElement().getChildNodes().item(1).getFirstChild().getAttributes().getNamedItem("style").getTextContent();

                  String pattern = ".*background-image: url\\('.*/(.*).png'\\);.*";

                  String icon = style.replaceAll(pattern, "$1");

                  if (runeType == null)
                  {
                     writer.write("SKILL_MAP.put(\"" + activeSlug + "\", new Skill(\"" + activeSlug + "\", \"" + icon + "\", \"" + skillName + "\"));\n");
                  }
                  else
                  {
                     String runeSlug = activeSlug + "-" + runeType;

                     String runeName = doc.getDocumentElement().getChildNodes().item(2).getChildNodes().item(1).getTextContent();
                     String runeText = doc.getDocumentElement().getChildNodes().item(2).getChildNodes().item(2).getTextContent();

                     writer.write("RUNE_MAP.put(\"" + runeSlug + "\", new Rune(\"" + runeSlug + "\", \"" + runeName + "\", \"" + runeText + "\"));\n");
                  }
               }
               catch (Exception ex)
               {
                  ex.printStackTrace();
                  throw ex;
               }
               finally
               {
                  response.close();
               }
            }
         }
      }

      writer.close();
   }
}
